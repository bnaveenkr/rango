package rango

import "math"

func Cone(object *Object, material Material, radius float64, height float64, resolution uint) {

	top := V(0, height, 0)
	origin := V(0, 0, 0)

	triangles := make([]Triangle, 0)

	var t1 float64
	var t2 float64
	var v0 Vector
	var v1 Vector

	for i := 0; i < int(resolution); i++ {

		t1 = float64(i) / float64(resolution)
		t1 = deg2rad(t1 * 360)

		t2 = float64(i+1) / float64(resolution)
		t2 = deg2rad(t2 * 360)

		v0 = V(radius*math.Cos(t1), 0, radius*math.Sin(t1))
		v1 = V(radius*math.Cos(t2), 0, radius*math.Sin(t2))

		triangles = append(triangles, Triangle{top, v1, v0})
		triangles = append(triangles, Triangle{v1, origin, v0})

	}

	SetObject(object, material, uint32(2*resolution), triangles)
}
