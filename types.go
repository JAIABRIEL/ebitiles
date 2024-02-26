// Package tilemap docstring
package tilemap

// Quad is a basic quad tree interface
type Quad interface {
	Redraw() TileImage
	GetTile(int, int) *Tile
	DrawRow3D(TileImage, int, int)
	Draw(TileImage)
	GiveBirth(int)
}

// Image is a placeholder for your own draw function
type TileImage interface {
	// DrawAt draws another image at a position on this image.
	// Be aware that position is given in tiles.
	DrawAt(TileImage, int, int)

	// New resets this image to a new image of a given size.
	New(int)
}
