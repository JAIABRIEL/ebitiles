package ebitiles

import (
	"github.com/JAIABRIEL/gonimator"
	"github.com/hajimehoshi/ebiten/v2"
)

type Layer struct {
	AnimationPlayer *gonimator.AnimationPlayer[int]
	IsAnimated      bool
	DrawPosX        float64
	DrawPosY        float64

	Animation []*ebiten.Image
	Image     *ebiten.Image
}

func (l *Layer) Draw(img *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(l.DrawPosX, l.DrawPosY)

	if !l.IsAnimated {
		img.DrawImage(l.Image, op)
	} else {
		img.DrawImage(l.Animation[l.AnimationPlayer.Get()], op)
	}
}

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
