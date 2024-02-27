package tilemap

import "github.com/hajimehoshi/ebiten/v2"

type Tile struct {
	Layers   []*ebiten.Image
	layers3d []*Layer3D
	Chunk
}

// Draw docstring
func (t *Tile) Draw(img *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(t.GlobalX), float64(t.GlobalY))
	img.DrawImage(t.buffered, op)
}

// Redraw docstring
func (t *Tile) Redraw() *ebiten.Image {
	if t.buffered == nil {
		t.buffered = ebiten.NewImage(t.tileSize, t.tileSize)
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
	t.Layers[layer] = img
}
