package rango

import "math"

const EPSILON float64 = 0.00001

type Vector struct {
	X float64
	Y float64
	Z float64
}

func V(x float64, y float64, z float64) Vector {
	return Vector{x, y, z }
}

func VecVecMult(a Vector, b Vector) Vector {
	return V(a.X * b.X, a.Y * b.Y, a.Z * b.Z)
}

func FloatVecMult(a float64, b Vector) Vector {
	return V(a * b.X, a * b.Y, a * b.Z)
}

func Dot(a Vector, b Vector) float64  {
	return a.X * b.X + a.Y * b.Y + a.Z * b.Z
}

func Cross(a Vector, b Vector) Vector {
	var x float64 = a.Y * b.Z - b.Y * a.Z
	var y float64 = a.Z * b.X - b.Z * a.X
	var z float64 = a.X * b.Y - b.X * a.Y

	return V(x, y, z)
}

func Add(a Vector, b Vector) Vector {
	return V(a.X + b.X, a.Y + b.Y, a.Z + b.Z)
}

func Subtract(a Vector, b Vector) Vector {
	return V(a.X - b.X, a.Y - b.Y, a.Z - b.Z)
}

func Negate(v Vector) Vector {
	return V(-v.X, -v.Y, -v.Z)
}

func Length(v Vector) float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y + v.Z * v.Z)
}

func Normalize(v Vector) Vector {

	var length = Length(v)
	if math.Abs(length) < EPSILON {
		return v
	}

	return V(v.X / length, v.Y / length, v.Z / length)
}

