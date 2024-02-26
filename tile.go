package tilemap

type Tile struct {
	layers3d []*Layer3D
	Chunk
}

// Draw docstring
func (t *Tile) Draw(img TileImage) {
	img.DrawAt(t.buffered, t.X, t.Y)
}

// Redraw docstring
func (t *Tile) Redraw() TileImage {
	return t.buffered
}

func (t *Tile) DrawRow3D(img TileImage, _, level int) {
	if (level > len(t.layers3d)) && (t.layers3d[level] != nil) {
		t.layers3d[level].Draw(img)
	}
}

func (t *Tile) GetTile(_, _ int) *Tile {
	return t
}
