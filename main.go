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
	theta float64
	x     float64
	y     float64
}

func (g *Game) Update() error {
	g.theta++
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
	// 画像サイズを取得
	b := img.Bounds()
	op := &ebiten.DrawImageOptions{}
	// 画像の中心を原点にして回転
	op.GeoM.Translate(-float64(b.Dx())/2, -float64(b.Dy())/2)
	op.GeoM.Rotate(g.theta * 2 * 3.14159 / 360)
	// x,yを画像の中心に移動
	op.GeoM.Translate(float64(b.Dx())/2, float64(b.Dy())/2)
	// そのあと移動
	op.GeoM.Translate(g.x, g.y)
	op.GeoM.Scale(2.0, 1.0)
	screen.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})
	screen.DrawImage(img, op)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Hello, World! %f", g.theta))
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
