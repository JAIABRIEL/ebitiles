// Package tilemap docstring
package tilemap

import "github.com/hajimehoshi/ebiten/v2"

// Quad is a basic quad tree interface
type Quad interface {
	Redraw() *ebiten.Image
	GetTile(int, int) *Tile
	DrawRow3D(*ebiten.Image, int, int)
	Draw(*ebiten.Image)
	GetByLevel(int, int, int) Quad
	Create(ChunkLevel, int, int, int, int, int)
	InsertTile(*ebiten.Image, int, int, int)
}

// Image is a placeholder for your own draw function
type TileImage interface {
	// DrawAt draws another image at a position on this image.
	// Be aware that position is given in tiles.
	DrawAt(TileImage, int, int)
	Clear()
	Init(int, int, int)
}
