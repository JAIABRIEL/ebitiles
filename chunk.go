package tilemap

type Chunk struct {
	quads [4]Quad
	X     int
	Y     int

	Size     int
	TileSize int

	sizeHalf int
	buffered Image
}

func (c *Chunk) DrawRow(r int, img TileImage) TileImage {
	return nil
}

func (c *Chunk) GetTile(x, y int) *Tile {
	return c.quads[c.posToIndex(x, y)].GetTile(
		x-(x/c.sizeHalf*c.sizeHalf),
		y-(y/c.sizeHalf*c.sizeHalf))
}

func (c *Chunk) posToIndex(x, y int) int {
	return (x/c.sizeHalf)%2 + ((y/c.sizeHalf)%2)*2
}

func (c *Chunk) Redraw() TileImage {
	buffered := c.quads[0].Redraw()

	return buffered
}
