package ball

import (
	"ballSimulator/math"
	"ballSimulator/rectangle"
	"ballSimulator/time"
	"errors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type Ball struct {
	Position math.Vector
	drag     float64
	mass     float64

	velocity math.Vector
	radius   float64

	ground *rectangle.Rectangle

	Time *time.Time
}

type NewBallOpts struct {
	Y    float64
	X    float64
	Mass float64
	Drag float64
}

func NewBall(groundPointer *rectangle.Rectangle, timePointer *time.Time, opts NewBallOpts) (*Ball, error) {

	if opts.Mass == 0 {
		opts.Mass = 1
	}

	if groundPointer == nil {
		return nil, errors.New("groundPointer is nil")
	}

	if timePointer == nil {
		return nil, errors.New("null Time Pointer")
	}

	return &Ball{
		Position: math.Vector{X: opts.X, Y: opts.Y},
		Time:     timePointer,
		mass:     opts.Mass,
		drag:     opts.Drag,
		ground:   groundPointer,
		radius:   75,
	}, nil
}

func getBallImage() *ebiten.Image {
	ballImg, _, err := ebitenutil.NewImageFromFile("ball.png")

	if err != nil {
		log.Fatal(err)
	}

	return ballImg
}

func (b *Ball) Draw(screen *ebiten.Image) {
	img := getBallImage()
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(0.2, 0.2)
	opts.GeoM.Translate(b.Position.X, b.Position.Y)
	screen.DrawImage(img, opts)
}

func (b *Ball) getDragForce() math.Vector {
	velocitySquared := b.velocity.Pow(2)
	dragForce := velocitySquared.MulScalar(-b.drag)
	return dragForce
}

func (b *Ball) AddForce(force math.Vector) {
	deltaTime := b.Time.DeltaTime

	accelerationProduct := force.Divide(b.mass)

	velocityProduct := accelerationProduct.MulScalar(deltaTime)

	b.velocity.X += velocityProduct.X
	b.velocity.Y += velocityProduct.Y
}

func (b *Ball) ApplyGravity() {
	b.AddForce(math.Vector{X: 0, Y: b.mass * 10})
}

func (b *Ball) updateVelocity() {
	b.Position = b.Position.Sum(b.velocity)
}

func (b *Ball) collideWithGround() bool {

	closestX := max(b.ground.Position.X, min(b.Position.X, b.ground.Position.X+float64(b.ground.Width)))
	closestY := max(b.ground.Position.Y, min(b.Position.Y, b.ground.Position.Y+float64(b.ground.Height)))

	return b.collideWithPoint(math.Vector{X: closestX, Y: closestY})

}

func (b *Ball) reflect() {
	b.velocity.X = -b.velocity.X
	b.velocity.Y = -b.velocity.Y
}

func (b *Ball) collideWithPoint(point math.Vector) bool {
	distance := b.Position.EuclideanDistance(point)

	return distance < b.radius
}

func (b *Ball) Update() {
	b.ApplyGravity()

	if b.collideWithGround() {
		b.reflect()
	}

	b.updateVelocity()
	b.AddForce(b.getDragForce())
}
