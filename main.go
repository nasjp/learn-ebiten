package main

import (
	"bytes"
	"image"
	"io"
	"log"

	_ "embed"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed gopher.png
	gopherImageBytes []byte
	gopherImage      = MustDecode(bytes.NewBuffer(gopherImageBytes))
	tmpImage         = ebiten.NewImage(320, 240)
)

func MustDecode(r io.Reader) *ebiten.Image {
	img, _, err := image.Decode(r)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

type Game struct {
	angle int
}

func (g *Game) Update() error {
	g.angle++
	g.angle %= 360

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	const scale = 8.0
	tmpImage.Clear()

	op1 := &ebiten.DrawImageOptions{}
	op1.GeoM.Scale(1/scale, 1/scale)

	tmpImage.DrawImage(gopherImage, op1)

	op2 := &ebiten.DrawImageOptions{}
	op2.GeoM.Scale(scale, scale)

	screen.DrawImage(tmpImage, op2)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	game := &Game{}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Tmp")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
