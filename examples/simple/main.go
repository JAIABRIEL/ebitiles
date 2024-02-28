package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/JAIABRIEL/ebitiles"
)

const (
	screenWidth  = 240
	screenHeight = 240
)

type Game struct {
	Map *ebitiles.TileMap
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Each tile will draw itself on the screen.
	g.Map.Draw(screen)
}

func (g *Game) Update() error {
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
	tm.Create(6, 32, 1)

	grass, err := getImage("Grass.png")
	if err != nil {
		fmt.Println("Something went wrong...", err)
		return
	}

	for x := 0; x < 32; x++ {
		for y := 0; y < 32; y++ {
			fmt.Println("insert", x, y)
			tm.InsertTile(grass, x, y, 0)
		}
	}

	tm.Redraw()

	g := &Game{
		Map: tm,
	}
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Simple Demo")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
