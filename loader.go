package ebitiles

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Loader struct {
	// TileMap represents the TileMap from which chunks are provided.
	TileMap *TileMap
	// ChunkPos represents the position at specified level.
	// e.g. GlobalPos / (ChunkSize * TileSize)
	ChunkPosX int
	ChunkPosY int
	Level     ChunkLevel
	Radius    int

	chunkSize int
}

func (l *Loader) Init() {
	l.chunkSize = int(math.Pow(2, float64(l.Level)))
}

func (l *Loader) Update(x, y int) {
	l.ChunkPosX = x / l.chunkSize
	l.ChunkPosY = y / l.chunkSize
}

func (l *Loader) Draw(img *ebiten.Image) {
	x := l.ChunkPosX - l.Radius

	for x <= l.ChunkPosX {
		y := l.ChunkPosY - l.Radius
		for y <= l.ChunkPosY {
			l.TileMap.TestQuadByLevel(l.Level, x, y).Draw(img)
			y++
		}
		x++
	}
}
