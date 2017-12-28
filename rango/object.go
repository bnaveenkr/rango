package rango

type Object struct {
	Material  Material   /* Material properties of this object */
	Ntris     uint64     /* Number of Triangles */
	Triangles []Triangle /* Triangle list of the object */
}

func SetObject(object *Object, material Material, ntris uint32, triangles []Triangle) *Object {

	object.Material = material
	object.Ntris = uint64(ntris)
	object.Triangles = triangles

	return object
}
