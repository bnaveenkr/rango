package rango

func PlaneXZ(object *Object, material Material, size float64) {

	side := size * 0.5
	triangles := make([]Triangle, 0)
	var v0 Vector
	var v1 Vector
	var v2 Vector

	v0 = V(side, 0, side)
	v1 = V(side, 0, -side)
	v2 = V(-side, 0, -side)
	triangles = append(triangles, Triangle{v0, v1, v2})

	v0 = V(side, 0, side)
	v1 = V(-side, 0, side)
	v2 = V(-side, 0, -side)
	triangles = append(triangles, Triangle{v2, v1, v0})

	SetObject(object, material, 2, triangles)

}
