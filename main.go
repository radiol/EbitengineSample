package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var img *ebiten.Image

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("gopher.png")
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	count int
	x     float64
	y     float64
}

func (g *Game) Update() error {
	g.count++
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.y--
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.y++
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.x--
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.x++
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.x, g.y)
	op.GeoM.Scale(1.5, 1.0)
	screen.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})
	screen.DrawImage(img, op)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Hello, World! %d", g.count))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
