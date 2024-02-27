// Package tilemap docstring
package tilemap

// Quad is a basic quad tree interface
type Quad interface {
	Redraw() *EbitImage
	GetTile(int, int) *Tile
	DrawRow3D(*EbitImage, int, int)
	Draw(*EbitImage)
	GetByLevel(int, int, int) Quad
	init(int)
}

// Image is a placeholder for your own draw function
type TileImage interface {
	// DrawAt draws another image at a position on this image.
	// Be aware that position is given in tiles.
	DrawAt(TileImage, int, int)
	Clear()
	Init(int, int, int)
}
