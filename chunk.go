package ebitiles

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Chunk contains multiple Quads.
// These might be either other chunks or tiles at level 0.
type Chunk struct {
	quads    [4]Quad
	buffered *ebiten.Image

	tileSize int

	isActive bool
	Translate
}

// InsertTile inserts an image at given position and layer.
// Calling this on Chunk works with relative positions only.
func (c *Chunk) InsertTile(img *ebiten.Image, x, y, layer int) {
	c.isActive = true
	c.quads[c.posToIndex(x, y)].InsertTile(
		img,
		c.translatePos(x),
		c.translatePos(y),
		layer)
}

// DrawRow3D is still WIP and will be implemented when 3D is introducted
func (c *Chunk) DrawRow3D(img *ebiten.Image, row, level int) {
	pos := c.posToIndex(0, row)
	c.quads[pos].DrawRow3D(img, c.translatePos(row), level)
	c.quads[pos+1].DrawRow3D(img, c.translatePos(row), level)
}

// getTiles returns a tile from a relative position.
// Only the TileMap takes global position.
func (c *Chunk) getTile(x, y int) *Tile {
	return c.quads[c.posToIndex(x, y)].getTile(
		x-(x/c.sizeHalf*c.sizeHalf),
		y-(y/c.sizeHalf*c.sizeHalf))
}

// Redraw completely recreates this chunk.
// Inserted tiles still remain but will redraw themselves from layers.
// Caution: This might take a few seconds and shoudn't be called in ticks.
// Use loading screens instead or call Redraw on low leveled chunks.
func (c *Chunk) Redraw() *ebiten.Image {
	if !c.isActive {
		return c.buffered
	}
	c.buffered.Clear()

	for i := range c.quads {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64((i%2)*c.sizeHalf), float64((i/2)*c.sizeHalf))
		c.buffered.DrawImage(c.quads[i].Redraw(), op)
	}

	return c.buffered
}

// Draw passes the image down to the tiles so they can draw themselves at their global position.
// This might be used for ebiten screen in the games Draw function.
func (c *Chunk) Draw(img *ebiten.Image) {
	if !c.isActive {
		return
	}
	for _, c := range c.quads {
		c.Draw(img)
	}
}

// GetByLevel walks down the levels and returns the current item
// when ChunkLevel reaches 0.
func (c *Chunk) GetByLevel(chunkLevel ChunkLevel, x, y int) Quad {
	if chunkLevel == 0 {
		return c
	} else {
		return c.quads[c.posToIndex(x, y)].GetByLevel(
			chunkLevel-1,
			c.translatePos(x),
			c.translatePos(y))
	}
}

// Create initializes this Quad, and also initializes its children
// until the ChunkLevel reaches 0.
// This method needs to be called to work with this chunk at all.
func (c *Chunk) Create(chunkLevel ChunkLevel, size, tileSize, layerAmount, globalX, globalY int) {
	c.buffered = ebiten.NewImage(size*tileSize, size*tileSize)
	c.sizeHalf = size / 2
	c.Size = size
	c.tileSize = tileSize
	c.GlobalX = globalX
	c.GlobalY = globalY

	for i := range c.quads {
		if chunkLevel > 1 {
			c.quads[i] = &Chunk{}
		} else {
			c.quads[i] = &Tile{}
		}
		c.quads[i].Create(chunkLevel-1,
			c.sizeHalf,
			tileSize,
			layerAmount,
			globalX+(i%2)*c.sizeHalf,
			globalY+(i/2)*c.sizeHalf)
	}
}

// SetActive enables/disables this chunk.
// It won't be drawn on given images and wont be redrawn.
// Children won't be drawn as well but still keep their current state.
func (c *Chunk) SetActive(v bool) {
	c.isActive = v
}
