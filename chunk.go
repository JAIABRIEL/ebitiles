package tilemap

import (
	"math"
)

type Chunk struct {
	quads    [4]Quad
	buffered TileImage

	X int
	Y int

	Size     int
	TileSize int

	sizeHalf int
}

func (c *Chunk) DrawRow3D(img TileImage, row, level int) {
	pos := c.posToIndex(0, row)
	c.quads[pos].DrawRow3D(img, row-(row/c.sizeHalf*c.sizeHalf), level)
	c.quads[pos+1].DrawRow3D(img, row-(row/c.sizeHalf*c.sizeHalf), level)
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
	c.buffered.New(c.Size)
	c.buffered.DrawAt(c.quads[0].Redraw(), 0, 0)
	c.buffered.DrawAt(c.quads[1].Redraw(), c.sizeHalf, 0)
	c.buffered.DrawAt(c.quads[2].Redraw(), 0, c.sizeHalf)
	c.buffered.DrawAt(c.quads[3].Redraw(), c.sizeHalf, c.sizeHalf)

	return c.buffered
}

func (c *Chunk) Draw(img TileImage) {
	for _, c := range c.quads {
		c.Draw(img)
	}
}

func (c *Chunk) GiveBirth(chunkLevel int) {
	if chunkLevel > 1 {
		for i := range c.quads {
			c.quads[i] = &Chunk{
				TileSize: c.TileSize,
				Size:     c.sizeHalf,
				sizeHalf: c.sizeHalf / 2,
			}
			c.quads[i].GiveBirth(chunkLevel - 1)
		}
	} else {
		for i := range c.quads {
			c.quads[i] = &Tile{}
		}
	}
}

func NewChunk(chunkLevel int) *Chunk {
	c := &Chunk{
		Size: int(math.Pow(2, 10)),
	}
	c.sizeHalf = c.Size / 2
	c.GiveBirth(chunkLevel - 1)

	return c
}
