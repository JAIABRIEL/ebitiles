package ebitiles

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Tile implements quad and is always on the lowest level.
// It represents a single tile on the map.
type Tile struct {
	Layers   []*ebiten.Image
	layers3d []*Layer3D
	Chunk
}

// Draw draws this tiles buffered image on a image based on global position.
func (t *Tile) Draw(img *ebiten.Image) {
	if t.isActive {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(t.GlobalX*t.tileSize), float64(t.GlobalY*t.tileSize))
		img.DrawImage(t.buffered, op)
	}
}

// Redraw places all of this tiles layers.
// result is saved in buffer.
func (t *Tile) Redraw() *ebiten.Image {
	if !t.isActive {
		return t.buffered
	}
	t.buffered.Clear()

	for _, l := range t.Layers {
		if l != nil {
			t.buffered.DrawImage(l, &ebiten.DrawImageOptions{})
		}
	}
	return t.buffered
}

// DrawRow3D is WIP and doesn't do anything for now.
func (t *Tile) DrawRow3D(img *ebiten.Image, _, level int) {
	if (level > len(t.layers3d)) && (t.layers3d[level] != nil) {
		t.layers3d[level].Draw(img)
	}
}

func (t *Tile) getTile(_, _ int) *Tile {
	return t
}

// GetByLevel will return this Tile, since tiles are always on level 0.
func (t *Tile) GetByLevel(_ ChunkLevel, _, _ int) Quad {
	return t
}

// Create sets this tiles parameters.
func (t *Tile) Create(_ ChunkLevel, size, tileSize, layerAmount, globalX, globalY int) {
	t.Size = size
	t.tileSize = tileSize
	t.GlobalX = globalX
	t.GlobalY = globalY

	t.buffered = ebiten.NewImage(tileSize, tileSize)
	t.Layers = make([]*ebiten.Image, layerAmount)
}

// InsertTile will set an *ebiten.Image on one of this tiles layers.
func (t *Tile) InsertTile(img *ebiten.Image, _, _, layer int) {
	t.isActive = true
	t.Layers[layer] = img
}
