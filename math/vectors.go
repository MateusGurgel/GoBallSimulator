package math

import "math"

type Vector struct {
	X float64
	Y float64
}

func (v Vector) Sum(vector Vector) Vector {
	return Vector{X: v.X + vector.X, Y: v.Y + vector.Y}
}

func (v Vector) MulScalar(scalar float64) Vector {
	return Vector{X: v.X * scalar, Y: v.Y * scalar}
}

func (v Vector) Pow(scalar float64) Vector {
	return Vector{X: math.Pow(v.X, scalar), Y: math.Pow(v.Y, scalar)}
}

func (v Vector) EuclideanDistance(vector Vector) float64 {
	return math.Sqrt(math.Pow(v.X-vector.X, 2) + math.Pow(v.Y-vector.Y, 2))
}

func (v Vector) Divide(scalar float64) Vector {
	return Vector{X: v.X / scalar, Y: v.Y / scalar}
}

func (v Vector) AddScalar(scalar float64) Vector {
	return Vector{X: v.X + scalar, Y: v.Y + scalar}
}

func (v Vector) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
