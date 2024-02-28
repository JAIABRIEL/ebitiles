package tilemap

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// ChunkLevel alias for uint16
type ChunkLevel = uint16

// TileMap should always on top of all quad items, since it translates
// all functions so they can be used from outside of this library.
// While all quad items use relative positive-only positions,
// TileMap works with both, positive and negative positions.
// It does not implement Quad.
type TileMap struct {
	quads [4]Quad
	level ChunkLevel
	Translate
}

// InsertTile inserts an image to this map on a global position and layer.
// Tiles are based on all layers placed on each other from lowest layer.
// Be aware that layer is an array index.
func (tm *TileMap) InsertTile(img *ebiten.Image, x, y, layer int) {
	tm.quads[tm.posToIndex(tm.toPositive(x), tm.toPositive(y))].InsertTile(
		img,
		tm.translateNegativePos(x),
		tm.translateNegativePos(y),
		layer)
}

// Create inisializes this TileMap with given parameters.
// This maps size is 2^level based on the level parameter.
// level must be >= 2
func (tm *TileMap) Create(level ChunkLevel, tileSize int, layerAmount int) {
	tm.Size = int(math.Pow(2, float64(level)))
	tm.sizeHalf = tm.Size / 2
	tm.level = level

	for i := range tm.quads {
		tm.quads[i] = &Chunk{}

		tm.quads[i].Create(level-1,
			tm.sizeHalf,
			tileSize,
			layerAmount,
			-tm.sizeHalf+(i%2)*tm.sizeHalf,
			-tm.sizeHalf+(i/2)*tm.sizeHalf)
	}
}

// Redraw buffers tiles and chunks.
// It should be called after the tiles are placed.
// Caution: Don't use this in your ebiten Draw function.
// On large maps, it might take its time.
// Use it in loading screens instead or call Redraw on smaller chunks or tiles.
func (tm *TileMap) Redraw() {
	for _, q := range tm.quads {
		q.Redraw()
	}
}

// GetQuadByLevel returns the quad from given position and level.
// For example:
// GetQuadByLevel(1, 0, 0) => Chunk starting at (0/0) containing 4 tiles
// GetQuadByLevel(0, 0, 0) => Tile on (0/0) as Quad
func (tm *TileMap) GetQuadByLevel(globalLevel ChunkLevel, x, y int) Quad {
	return tm.quads[tm.posToIndex(tm.toPositive(x), tm.toPositive(y))].GetByLevel(
		tm.level-(globalLevel%tm.level),
		tm.translatePos(x),
		tm.translatePos(y))
}

// Draw will pass an *ebiten.Image down the tree.
// Each tile will draw itself on the image.
// Call this in your games Draw function.
// Empty chunks and tiles are skipped.
func (tm *TileMap) Draw(img *ebiten.Image) {
	for _, q := range tm.quads {
		q.Draw(img)
	}
}
