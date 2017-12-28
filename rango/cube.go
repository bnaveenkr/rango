package rango

func Cube(object *Object, material Material, side float64) {

	var triangles = make([]Triangle, 0)

	var halfSide float64 = side * 0.5

	min := V(-halfSide, -halfSide, -halfSide)
	max := V(halfSide, halfSide, halfSide)

	var v0 Vector = Vector{}
	var v1 Vector = Vector{}
	var v2 Vector = Vector{}

	/* front face */
	v0 = V(min.X, min.Y, min.Z)
	v1 = V(max.X, min.Y, min.Z)
	v2 = V(max.X, max.Y, min.Z)
	triangles = append(triangles, Triangle{v2, v1, v0})

	v0 = V(max.X, max.Y, min.Z)
	v1 = V(min.X, max.Y, min.Z)
	v2 = V(max.X, min.Y, min.Z)
	triangles = append(triangles, Triangle{v2, v1, v0})

	/* right */
	v0 = V(max.X, min.Y, min.Z)
	v1 = V(max.X, min.Y, max.Z)
	v2 = V(max.X, max.Y, max.Z)
	triangles = append(triangles, Triangle{v2, v1, v0})

	v0 = V(max.X, max.Y, max.Z)
	v1 = V(max.X, max.Y, min.Z)
	v2 = V(max.X, min.Y, min.Z)
	triangles = append(triangles, Triangle{v2, v1, v0})

	/* back */
	v0 = V(max.X, min.Y, max.Z)
	v1 = V(min.X, min.Y, max.Z)
	v2 = V(min.X, max.Y, max.Z)
	triangles = append(triangles, Triangle{v2, v1, v0})

	v0 = V(min.X, max.Y, max.Z)
	v1 = V(max.X, max.Y, max.Z)
	v2 = V(max.X, min.Y, max.Z)
	triangles = append(triangles, Triangle{v2, v1, v0})

	/* left */
	v0 = V(min.X, min.Y, max.Z)
	v1 = V(min.X, min.Y, min.Z)
	v2 = V(min.X, max.Y, min.Z)
	triangles = append(triangles, Triangle{v2, v1, v0})

	v0 = V(min.X, max.Y, min.Z)
	v1 = V(min.X, max.Y, max.Z)
	v2 = V(min.X, min.Y, max.Z)

	triangles = append(triangles, Triangle{v2, v1, v0})

	/* bottom */
	v0 = V(min.X, min.Y, min.Z)
	v1 = V(min.X, min.Y, max.Z)
	v2 = V(max.X, min.Y, max.Z)

	triangles = append(triangles, Triangle{v2, v1, v0})

	v0 = V(max.X, min.Y, max.Z)
	v1 = V(max.X, min.Y, min.Z)
	v2 = V(min.X, min.Y, min.Z)
	triangles = append(triangles, Triangle{v2, v1, v0})

	/* top */
	v0 = V(min.X, max.Y, min.Z)
	v1 = V(max.X, max.Y, min.Z)
	v2 = V(max.X, max.Y, max.Z)
	triangles = append(triangles, Triangle{v2, v1, v0})

	v0 = V(max.X, max.Y, max.Z)
	v1 = V(min.X, max.Y, max.Z)
	v2 = V(min.X, max.Y, min.Z)
	triangles = append(triangles, Triangle{v2, v1, v0})

	SetObject(object, material, 12, triangles)
}
