package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/JAIABRIEL/gonimator"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/JAIABRIEL/ebitiles"
)

const (
	screenWidth  = 240
	screenHeight = 240
)

var WaterAssets []*ebiten.Image

var WaterAnimation = &gonimator.Animation[int]{
	Parts: []*gonimator.Part[int]{
		{Value: 0, Duration: 10},
		{Value: 1, Duration: 10},
		{Value: 2, Duration: 10},
		{Value: 3, Duration: 10},
		{Value: 4, Duration: 10},
		{Value: 5, Duration: 10},
		{Value: 6, Duration: 10},
		{Value: 7, Duration: 10},
	},
}

type Game struct {
	Map    *ebitiles.TileMap
	Player *gonimator.AnimationPlayer[int]
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Each tile will draw itself on the screen.
	g.Map.Draw(screen)
}

func (g *Game) Update() error {
	g.Player.Update()
	fmt.Println(ebiten.ActualFPS())
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func getImage(imgPath string) (*ebiten.Image, error) {
	file, err := os.ReadFile(imgPath)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(bytes.NewReader(file))
	if err != nil {
		return nil, err
	}

	return ebiten.NewImageFromImage(img), nil
}

func main() {
	tm := &ebitiles.TileMap{}
	tm.Create(6, 64, 2)

	assets := loadAssets()
	fmt.Println(assets)
	ap := WaterAnimation.NewPlayer()

	for x := 0; x < 32; x++ {
		for y := 0; y < 32; y++ {
			tm.InsertTileAnimated(assets, ap, 1, x/2, y)
		}
	}

	tm.Redraw()

	g := &Game{
		Map:    tm,
		Player: ap,
	}
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Simple Demo")
	ebiten.SetVsyncEnabled(false)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func loadAssets() []*ebiten.Image {
	files := []string{
		"WaterBase1.png",
		"WaterBase2.png",
		"WaterBase3.png",
		"WaterBase4.png",
		"WaterBase5.png",
		"WaterBase6.png",
		"WaterBase7.png",
		"WaterBase8.png",
	}

	assets := make([]*ebiten.Image, len(files))

	for i, f := range files {
		assets[i], _ = getImage("assets/BetterWater/" + f)
	}
	return assets
}
