package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/JAIABRIEL/ebitiles"
)

const (
	screenWidth  = 400
	screenHeight = 400
)

type Game struct {
	Map    *ebitiles.TileMap
	Loader *ebitiles.Loader

	PlayerImg *ebiten.Image
	PosX      int
	PosY      int
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Each tile will draw itself on the screen.
	g.Loader.Draw(screen)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.PosX*32+8), float64(g.PosY*32+8))

	screen.DrawImage(g.PlayerImg, op)
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		g.PosY -= 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		g.PosX -= 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		g.PosY += 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		g.PosX += 1
	}

	g.Loader.Update(g.PosX, g.PosY)
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
	tm.Create(8, 32, 1)

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
		Map:    tm,
		Loader: tm.NewLoader(2, 0),
	}

	g.Loader.Update(0, 0)

	g.PlayerImg, _ = getImage("Player.png")

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Simple Demo")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
