package tilemap

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type ChunkLevel = uint16

type TileMap struct {
	quads [4]Quad
	level ChunkLevel
	Translate
}

func (tm *TileMap) InsertTile(img *ebiten.Image, x, y, layer int) {
	tm.quads[tm.posToIndex(tm.toPositive(x), tm.toPositive(y))].InsertTile(
		img,
		tm.translateNegativePos(x),
		tm.translateNegativePos(y),
		// x-(x/tm.sizeHalf*tm.sizeHalf),
		// y-(y/tm.sizeHalf),
		layer)
}

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

func (tm *TileMap) Redraw() {
	for _, q := range tm.quads {
		q.Redraw()
	}
}

func (tm *TileMap) GetChunkByLevel(x, y, level int) Quad {
	return tm.quads[tm.posToIndex(tm.toPositive(x), tm.toPositive(y))].GetByLevel(
		tm.translatePos(x),
		tm.translatePos(y),
		level)
}

func (tm *TileMap) Draw(img *ebiten.Image) {
	for _, q := range tm.quads {
		q.Draw(img)
	}
}
