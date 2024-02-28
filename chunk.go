package tilemap

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Chunk struct {
	quads    [4]Quad
	buffered *ebiten.Image

	// each chunk knows its own global position
	// for better performance
	GlobalX int
	GlobalY int

	tileSize int

	isActive bool
	Translate
}

func (c *Chunk) InsertTile(img *ebiten.Image, x, y, layer int) {
	c.isActive = true
	c.quads[c.posToIndex(x, y)].InsertTile(
		img,
		x-(x/c.sizeHalf*c.sizeHalf),
		y-(y/c.sizeHalf*c.sizeHalf),
		layer)
}

func (c *Chunk) DrawRow3D(img *ebiten.Image, row, level int) {
	pos := c.posToIndex(0, row)
	c.quads[pos].DrawRow3D(img, row-(row/c.sizeHalf*c.sizeHalf), level)
	c.quads[pos+1].DrawRow3D(img, row-(row/c.sizeHalf*c.sizeHalf), level)
}

func (c *Chunk) GetTile(x, y int) *Tile {
	return c.quads[c.posToIndex(x, y)].GetTile(
		x-(x/c.sizeHalf*c.sizeHalf),
		y-(y/c.sizeHalf*c.sizeHalf))
}

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

func (c *Chunk) Draw(img *ebiten.Image) {
	for _, c := range c.quads {
		c.Draw(img)
	}
}

func (c *Chunk) GetByLevel(x, y, chunkLevel int) Quad {
	if chunkLevel == 0 {
		return c
	} else {
		return c.quads[c.posToIndex(x, y)].GetByLevel(
			x-(x/c.sizeHalf*c.sizeHalf),
			y-(y/c.sizeHalf*c.sizeHalf),
			chunkLevel-1)
	}
}

func (c *Chunk) Create(chunkLevel ChunkLevel, size, tileSize, layerAmount, globalX, globalY int) {
	c.buffered = ebiten.NewImage(size*tileSize, size*tileSize)
	c.sizeHalf = size / 2
	c.Size = size
	c.tileSize = tileSize
	c.GlobalX = globalX
	c.GlobalY = globalY

	if chunkLevel > 1 {
		for i := range c.quads {
			c.quads[i] = &Chunk{}
			c.quads[i].Create(chunkLevel-1,
				c.sizeHalf,
				tileSize,
				layerAmount,
				globalX+(i%2)*c.sizeHalf,
				globalY+(i/2)*c.sizeHalf)
			// c.translateNegativePos(globalX),
			// c.translateNegativePos(globalY))
		}
	} else {
		for i := range c.quads {
			c.quads[i] = &Tile{}
			c.quads[i].Create(chunkLevel-1,
				c.sizeHalf,
				tileSize,
				layerAmount,
				globalX+(i%2)*c.sizeHalf,
				globalY+(i/2)*c.sizeHalf)

			// c.translatePos(globalX),
			// c.translatePos(globalY))
		}
	}
	// fmt.Println("pos", globalX, c.translatePos(globalX))
}
