package tilemap

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type ChunkLevel = uint16

type TileMap struct {
	quads    [4]Quad
	level    ChunkLevel
	size     int
	sizeHalf int
}

func (tm *TileMap) InsertTile(img *ebiten.Image, x, y, layer int) {
	fmt.Println("Insert Tile:", x, y, tm.translatePos(x), tm.translatePos(y))
	fmt.Println(tm.toPositive(x), tm.toPositive(y))
	fmt.Println(tm.posToIndex(tm.toPositive(x), tm.toPositive(y)))
	tm.quads[tm.posToIndex(tm.toPositive(x), tm.toPositive(y))].InsertTile(
		img,
		tm.translatePos(x),
		tm.translatePos(y),
		// x-(x/tm.sizeHalf*tm.sizeHalf),
		// y-(y/tm.sizeHalf),
		layer)
}

func (tm *TileMap) Create(level ChunkLevel, tileSize int, layerAmount int) {
	tm.size = int(math.Pow(2, float64(level)))
	tm.sizeHalf = tm.size / 2
	tm.level = level

	for i := range tm.quads {
		tm.quads[i] = &Chunk{
			Size: tm.size / 2,
		}

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

func (tm *TileMap) Draw(img *ebiten.Image) {
	for _, q := range tm.quads {
		q.Draw(img)
	}
}

func (tm *TileMap) translatePos(p int) int {
	p = tm.toPositive(p)
	return p - (p/tm.sizeHalf)*tm.sizeHalf
}

// toPositive translates the position to a positive tilemap starting from 0
func (tm *TileMap) toPositive(i int) int {
	return (i + tm.sizeHalf)
}

func (tm *TileMap) posToIndex(x, y int) int {
	// return (y%2)*2 + x%2
	return (x / tm.sizeHalf) + (y/tm.sizeHalf)*2
}
