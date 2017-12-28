package rango

func TransformObject(object *Object, M Matrix) {

	var v0, v1, v2 Vector

	for i := 0; i < int(object.Ntris); i++ {
		v0 = MatrixVecMultiply(M, object.Triangles[i].V0)
		v1 = MatrixVecMultiply(M, object.Triangles[i].V1)
		v2 = MatrixVecMultiply(M, object.Triangles[i].V2)
		object.Triangles[i] = Triangle{v0, v1, v2}
	}
}
