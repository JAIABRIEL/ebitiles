package tilemap

// Translate is a utility struct that implements methods that are
// required by other structs.
// It's made for TileMap so it doesn't have to implement the Chunk struct.
type Translate struct {
	GlobalX  int
	GlobalY  int
	Size     int
	sizeHalf int
}

// toPositive translates the position to a positive tilemap starting from 0
// for 10x10 grid: toPositive(1) => -6
func (t *Translate) toPositive(i int) int {
	return (i + t.sizeHalf)
}

// posToIndex returns an index position based on x y.
// It's made for [4]int to it will return 0-3.
func (tm *Translate) posToIndex(x, y int) int {
	return (x / tm.sizeHalf) + (y/tm.sizeHalf)*2
}

// translatePos translates a grid position to a position
// that is relative to a grid with half the size.
func (c *Translate) translatePos(p int) int {
	return p - p/c.sizeHalf*c.sizeHalf
}

// translateNegativePos can be used to receive a relative positive position
// from another position that allows negative positions.
// for 10x10 grid: translateNegativePos(-4) => 1
func (t *Translate) translateNegativePos(p int) int {
	return t.translatePos(t.toPositive(p))
}
