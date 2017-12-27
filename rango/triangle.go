package rango

type Triangle struct {
	V0, V1, V2 Vector
}

func SetTriangle(triangle *Triangle, v0 Vector, v1 Vector, v2 Vector) *Triangle {
	triangle.V0 = v0
	triangle.V1 = v1
	triangle.V2 = v2
	return triangle
}
