package rango

import "math"

func subdivide(src Triangle, dest *[4]Triangle, radius float64)  {
	/* mid points of the edges of the triangle */
	mid0 := FloatVecMult(radius, Normalize(Add(src.V0, src.V1)))
	mid1 := FloatVecMult(radius, Normalize(Add(src.V1, src.V2)))
	mid2 := FloatVecMult(radius, Normalize(Add(src.V2, src.V0)))

	/* input triangle subdivided into 4 based on the edges mid point */
	SetTriangle(&dest[0], mid0,   mid1,   mid2);
	SetTriangle(&dest[1], src.V0, mid0,   mid2);
	SetTriangle(&dest[2], mid0,   src.V1, mid1);
	SetTriangle(&dest[3], mid2,   mid1,   src.V2);
}

func subdivideRecursively(recursionDepth int, in Triangle, out []Triangle, outIndex int, radius float64)  {

	var triangles [4]Triangle
	/* span is the amount subdivide should jump so that final triangles are placed fine */
	span := int(math.Pow(4.0, float64(recursionDepth-1)))

	if recursionDepth > 0 {
		/* in triangle subdivided, and those are further sent for subdivision */
		subdivide(in, &triangles, radius)
		for i := 0; i<4; i++ {
			subdivideRecursively(recursionDepth-1, triangles[i], out, outIndex + i * span, radius)
		}
	} else {
		/* recursion has ended, lets populate the incoming triangle to *out */
		out[outIndex] = in
	}

}

func Sphere(object *Object, material Material, radius float64,resolution int) *Object {

	var octaHedron [8]Triangle
	var vertices [6]Vector
	triangleCount := 8 * math.Pow(4.0, float64(resolution))
	out := make([]Triangle, int(triangleCount), int(triangleCount))

	/* Sphere starts as an octahedron, minimum triangles is 8 and subdivides to 4 each time */
	vertices[0] = V(0.0, 0.0, -radius)
	vertices[1] = V(0.0, -radius, 0.0)
	vertices[2] = V(-radius, 0.0, 0.0)
	vertices[3] = V(0.0, 0.0,  radius)
	vertices[4] = V(0.0, radius, 0.0)
	vertices[5] = V(radius, 0.0, 0.0)

	/* Octahedron's top triangles */
	SetTriangle(&octaHedron[0], vertices[5], vertices[4], vertices[3])
	SetTriangle(&octaHedron[1], vertices[3], vertices[4], vertices[2])
	SetTriangle(&octaHedron[2], vertices[2], vertices[4], vertices[0])
	SetTriangle(&octaHedron[3], vertices[0], vertices[4], vertices[5])

	/* Octahedron's bottom triangles */
	SetTriangle(&octaHedron[4], vertices[3], vertices[1], vertices[5])
	SetTriangle(&octaHedron[5], vertices[5], vertices[1], vertices[0])
	SetTriangle(&octaHedron[6], vertices[0], vertices[1], vertices[2])
	SetTriangle(&octaHedron[7], vertices[2], vertices[1], vertices[2])

	/* Lets start subdividing these 8 triangles recursively */
	for i :=0; i<8; i++ {
		subdivideRecursively(resolution, octaHedron[i], out, i*int(triangleCount)/8, radius)
	}

	/* finally set the object */
	SetObject(object, material, uint32(len(out)), out)
	return object
}
