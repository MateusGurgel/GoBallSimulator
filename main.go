package main

import (
	"ballSimulator/ball"
	"ballSimulator/math"
	"ballSimulator/rectangle"
	"ballSimulator/time"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type Game struct {
	ground *rectangle.Rectangle
	ball   *ball.Ball
	time   *time.Time
}

func (g *Game) Update() error {
	g.time.UpdateDeltaTime()
	g.ball.Update()

	return nil
}

func (g *Game) init() {
	g.time = time.NewTime()
	g.ground = &rectangle.Rectangle{Position: math.Vector{X: 0.0, Y: 620}, Width: 640, Height: 20}
	g.ball, _ = ball.NewBall(g.ground, g.time, ball.NewBallOpts{X: 260.0, Y: 90.0, Drag: 0.1, Mass: 1})
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.ball.Draw(screen)
	g.ground.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 640
}

func main() {
	ebiten.SetWindowSize(640, 640)
	ebiten.SetWindowTitle("Hello, Ball!")

	game := Game{}

	game.init()

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}

}
