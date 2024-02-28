// Package tilemap docstring
package tilemap

import "github.com/hajimehoshi/ebiten/v2"

// Quad is a basic quad tree interface.
// It's implemented by Chunk and Tile.
type Quad interface {
	// Redraw completely recreates this quad and all of its children.
	// Inserted tiles still remain but will redraw themselves from layers.
	// Caution: This might take a few seconds and shoudn't be called in ticks.
	// Use loading screens instead or call Redraw on tiles or low leveled chunks.
	Redraw() *ebiten.Image

	// DrawRow3D does nothing so far since 3D layers are not implemented yet.
	DrawRow3D(*ebiten.Image, int, int)

	// Draw draws all of this quads children on a given *ebiten.Image
	// Might be used to draw on ebiten screen in the games draw function.
	Draw(*ebiten.Image)

	// GetByLevel does not return the actual level given.
	// Instead it walks down to the position and returns the item
	// as soon as the level parameter reaches 0
	// or tile level (level 0) is reached.
	GetByLevel(ChunkLevel, int, int) Quad

	// Create initializes this Quad, and also initializes its children
	// until the ChunkLevel reaches 0.
	// This method needs to be called to work with this Quad at all.
	Create(ChunkLevel, int, int, int, int, int)

	// InsertTiles inserts an *ebiten.Image  on a relative position and layer.
	// Image size should be equal to tile size.
	// Using different sized images might result in very shady results.
	InsertTile(*ebiten.Image, int, int, int)

	// SetActive enables/disables this Quad.
	// It won't be drawn on given images and wont be redrawn.
	// Children won't be drawn as well but still keep their current state.
	SetActive(bool)

	getTile(int, int) *Tile
}
