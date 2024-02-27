package tilemap

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Tile struct {
	Layers   []*ebiten.Image
	layers3d []*Layer3D
	Chunk

	isActive bool
}

// Draw docstring
func (t *Tile) Draw(img *ebiten.Image) {
	if t.isActive {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(t.GlobalX*t.tileSize), float64(t.GlobalY*t.tileSize))
		img.DrawImage(t.buffered, op)
	}
}

// Redraw docstring
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

func (t *Tile) DrawRow3D(img *ebiten.Image, _, level int) {
	if (level > len(t.layers3d)) && (t.layers3d[level] != nil) {
		t.layers3d[level].Draw(img)
	}
}

func (t *Tile) GetTile(_, _ int) *Tile {
	return t
}

func (t *Tile) GetByLevel(_, _, _ int) Quad {
	return t
}

func (t *Tile) Create(_ ChunkLevel, size, tileSize, layerAmount, globalX, globalY int) {
	t.Size = size
	t.tileSize = tileSize
	t.GlobalX = globalX
	t.GlobalY = globalY

	t.buffered = ebiten.NewImage(tileSize, tileSize)
	t.Layers = make([]*ebiten.Image, layerAmount)
}

func (t *Tile) InsertTile(img *ebiten.Image, x, y, layer int) {
	t.isActive = true
	t.Layers[layer] = img
}
