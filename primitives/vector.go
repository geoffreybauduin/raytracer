package primitives

import (
	"image/color"
	"math"
)

type Vector struct {
	X float64
	Y float64
	Z float64
}

func (v Vector) Subtract(oth Vector) Vector {
	return Vector{
		X: v.X - oth.X,
		Y: v.Y - oth.Y,
		Z: v.Z - oth.Z,
	}
}

func (v Vector) DotProduct(oth Vector) float64 {
	return v.X*oth.X + v.Y*oth.Y + v.Z*oth.Z
}

func (v Vector) Add(oth Vector) Vector {
	return Vector{
		X: v.X + oth.X,
		Y: v.Y + oth.Y,
		Z: v.Z + oth.Z,
	}
}

func (v Vector) Multiply(oth Vector) Vector {
	return Vector{
		X: v.X * oth.X,
		Y: v.Y * oth.Y,
		Z: v.Z * oth.Z,
	}
}

func (v Vector) MultiplyScalar(s float64) Vector {
	return Vector{
		X: v.X * s,
		Y: v.Y * s,
		Z: v.Z * s,
	}
}

func (v Vector) Divide(s float64) Vector {
	return Vector{
		X: v.X / s,
		Y: v.Y / s,
		Z: v.Z / s,
	}
}

func (v Vector) Length() float64 {
	return math.Sqrt(v.SquaredLength())
}

func (v Vector) SquaredLength() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vector) Unit() Vector {
	return v.Divide(v.Length())
}

var (
	Vec1 = Vector{1.0, 1.0, 1.0}
)

func (v Vector) ToColor() color.RGBA {
	return color.RGBA{
		R: uint8(255.99 * math.Sqrt(v.X)),
		G: uint8(255.99 * math.Sqrt(v.Y)),
		B: uint8(255.99 * math.Sqrt(v.Z)),
		A: 255,
	}
}
