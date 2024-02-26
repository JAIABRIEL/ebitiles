package tilemap

type Tile struct {
	layers3d []*Layer3D
	Chunk
}

// Draw docstring
func (t *Tile) Draw(img *EbitImage) {
	img.DrawAt(t.buffered, t.GlobalX, t.GlobalY)
}

// Redraw docstring
func (t *Tile) Redraw() *EbitImage {
	return t.buffered
}

func (t *Tile) DrawRow3D(img *EbitImage, _, level int) {
	if (level > len(t.layers3d)) && (t.layers3d[level] != nil) {
		t.layers3d[level].Draw(img)
	}
}

func (t *Tile) GetTile(_, _ int) *Tile {
	return t
}
