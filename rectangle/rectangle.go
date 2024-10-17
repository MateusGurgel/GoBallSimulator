package rectangle

import (
	"ballSimulator/math"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type Rectangle struct {
	Position math.Vector
	Width    int
	Height   int
}

func (r Rectangle) Draw(screen *ebiten.Image) {
	rectImage := ebiten.NewImage(r.Width, r.Height)
	rectImage.Fill(color.RGBA{G: 255, A: 255})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(r.Position.X, r.Position.Y)
	screen.DrawImage(rectImage, op)
}
