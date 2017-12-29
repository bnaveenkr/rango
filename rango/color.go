package rango

import "math"

const Uint8Max = 1<<8 - 1

type Color struct {
	R uint8
	G uint8
	B uint8
}

func Vector2Color(v Vector) Color {
	c := Color{}
	c.R = uint8(math.Min(1.0, math.Max(0.0, v.X)) * Uint8Max)
	c.G = uint8(math.Min(1.0, math.Max(0.0, v.Y)) * Uint8Max)
	c.B = uint8(math.Min(1.0, math.Max(0.0, v.Z)) * Uint8Max)
	return c
}

func MultiplyColor(a Color, b Color) Color {

	vector := Vector{}
	vector.X = float64(a.R * b.R)
	vector.Y = float64(a.G * b.G)
	vector.Z = float64(a.B * a.B)

	return Vector2Color(Normalize(vector))
}

func MultiplyColorVector(a Color, b Color) Vector {

	vector := Vector{}
	vector.X = float64(a.R * b.R)
	vector.Y = float64(a.G * b.G)
	vector.Z = float64(a.B * a.B)

	return vector
}
