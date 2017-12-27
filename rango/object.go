package rango

type Object struct {
	material  Material   /* Material properties of this object */
	ntris     uint64     /* Number of Triangles */
	triangles []Triangle /* Triangle list of the object */
}
