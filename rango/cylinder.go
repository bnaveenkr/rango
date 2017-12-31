package rango

import "math"

func Cylinder(object *Object, material Material, radius float64, height float64, resolution uint64) {

	top := V(0, height, 0)
	origin := V(0, 0, 0)

	triangles := make([]Triangle, 0)

	var v0, v1, v2, v3 Vector
	var t1, t2 float64

	for i := 0; i < int(resolution); i++ {
		t1 = float64(i) / float64(resolution)
		t1 = deg2rad(t1 * 360)

		t2 = float64(i+1) / float64(resolution)
		t2 = deg2rad(t2 * 360)

		v0 = V(radius*math.Cos(t1), 0, radius*math.Sin(t1))
		v1 = V(radius*math.Cos(t2), 0, radius*math.Sin(t2))
		v2 = V(v1.X, height, v1.Z)
		v3 = V(v0.X, height, v0.Z)

		/* Cylinder structure */
		triangles = append(triangles, Triangle{v2, v1, v0})
		triangles = append(triangles, Triangle{v0, v3, v2})

		/* Cylinder base */
		triangles = append(triangles, Triangle{v1, origin, v0})

		/* Cylinder top */
		triangles = append(triangles, Triangle{top, v2, v3})
	}

	SetObject(object, material, uint32(4*resolution), triangles)
}
