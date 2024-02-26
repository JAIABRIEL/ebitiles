package tilemap

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type EbitImage struct {
	Image    *ebiten.Image
	TileSize int
}

func (e *EbitImage) DrawAt(img *EbitImage, x, y int) {
	if img == nil {
		return
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x*e.TileSize), float64(x*e.TileSize))

	e.Image.DrawImage(img.Image, op)
}

func (e *EbitImage) Clear() {
	e.Image.Clear()
}

func (e *EbitImage) Init(width, height, tileSize int) {
	e.TileSize = tileSize
	e.Image = ebiten.NewImage(width, height)
}
