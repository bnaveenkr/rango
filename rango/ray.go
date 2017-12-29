package rango

import (
	"math"
)

const FARPLANE = 1 << 30

type Ray struct {
	Src Vector
	Dir Vector
}

type Hit struct {
	Position Vector  /* Position of the ray hit point */
	Normal   Vector  /* normal at that point */
	Ray      Ray     /* incident ray at the hit point */
	ObjectId int32  /* Object id of the hit object */
	T        float64 /* Distance from camera position to nearest hit */
}

func SetRay(ray *Ray, src Vector, dir Vector) {
	ray.Src = src
	ray.Dir = dir
}

func GenerateRay(i int, j int, cam Camera) Ray {

	samplei := (2*(float64(i)/float64(cam.Width)) - 1) * cam.Ar
	samplej := 2*(float64(j)/float64(cam.Height)) - 1

	sampleHori := FloatVecMult(samplei*cam.Fov2, cam.Hori)
	sampleVert := FloatVecMult(samplej*cam.Fov2, cam.Vert)
	sampleLook := Add(Add(sampleHori, sampleVert), cam.Look)

	ray := Ray{}
	SetRay(&ray, cam.Position, Normalize(sampleLook))

	return ray
}

func IntersectTriangle(ray Ray, triangle Triangle) float64 {

	var u, v, t, a, ia float64

	edge1 := Subtract(triangle.V1, triangle.V0)
	edge2 := Subtract(triangle.V2, triangle.V0)

	h := Cross(ray.Dir, edge2)
	a = Dot(edge1, h)

	if a < math.Abs(EPSILON) {
		return 0
	}

	ia = 1.0 / a

	s := Subtract(ray.Src, triangle.V0)
	u = Dot(s, h) * ia

	if u < 0 || u > 1 {
		return 0
	}

	q := Cross(s, edge1)
	v = Dot(ray.Dir, q) * ia
	if u < 0 || u+v > 1 {
		return 0
	}

	t = Dot(edge2, q) * ia
	if t < EPSILON {
		return 0
	}

	return t
}

func IntersectObject(ray Ray, object Object, objId int) Hit {

	var near float64 = FARPLANE
	var t float64
	var hitNormal Vector

	hit := Hit{}
	hit.T = 0

	for i := 0; i < int(object.Ntris); i++ {
		t = IntersectTriangle(ray, object.Triangles[i])

		/* Self occlusion check */
		if t > 0 && t < near {

			normal := Normalize(Cross(Subtract(object.Triangles[i].V1, object.Triangles[i].V0), Subtract(object.Triangles[i].V2, object.Triangles[i].V0)))

			/* only front facing triangles respond */
			if Dot(normal, ray.Dir) < EPSILON {
				near = t
				hitNormal = normal
			}
		}
	}

	/* near and far sanity check for hit */
	if near > 0 && near < FARPLANE {
		hit.Position = Add(ray.Src, FloatVecMult(near, ray.Dir))
		hit.Normal = hitNormal
		hit.ObjectId = int32(objId)
		hit.T = near
		return hit
	}

	/* objid = -1 means no hit, other vars may contain garbage values */
	hit.ObjectId = -1
	return hit
}

func IntersectScene(ray Ray, scene Scene) Hit {

	hit := Hit{}
	nearHit := Hit{}
	nearHit.T = FARPLANE

	for i := 0; i < int(scene.NObjects); i++ {
		hit = IntersectObject(ray, scene.Objects[i], i)

		/* Occlusion check */
		if hit.T > 0 && hit.T < nearHit.T {
			nearHit = hit
		}
	}
	/* near and far sanity check for hit */
	if nearHit.T > 0 && nearHit.T < FARPLANE {
		return nearHit
	}

	/* objid = -1 means no hit, other vars may contain garbage values */
	nearHit.ObjectId = -1
	return nearHit
}
