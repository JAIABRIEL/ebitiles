package tilemap

import "github.com/hajimehoshi/ebiten/v2"

type Layer3D struct {
	Image    *ebiten.Image
	GlobalX  int
	GlobalY  int
	tileSize int
}

func (l *Layer3D) Draw(img *ebiten.Image) {
	// op := &ebiten.DrawImageOptions{}
	// op.GeoM.Translate(float64(l.GlobalX*l.tileSize), float64(l.GlobalY*l.tileSize))
	// img.D
}
