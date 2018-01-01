package main

import (
	"fmt"
	"github.com/bnaveenkr/rango/rango"
)

func CreateScene(scene *rango.Scene) *rango.Scene {
	/* colors */
	red := rango.Color{255, 0, 0}
	green := rango.Color{0, 255, 0}
	blue := rango.Color{0, 0, 255}
	purple := rango.Color{255, 0, 255}
	white := rango.Color{255, 255, 255}

	/* Materials */
	mtlShiny1 := rango.Material{}
	rango.SetMaterial(&mtlShiny1, white, 0.1, 0.5, 0.4, 2.0, 0.4, 0.2, 1.4)

	mtlShiny2 := rango.Material{}
	rango.SetMaterial(&mtlShiny2, blue, 0.1, 0.5, 0.4, 3.0, 0.2, 0, 1.4)

	mtlMatte1 := rango.Material{}
	rango.SetMaterial(&mtlMatte1, red, 0.1, 0.5, 0.4, 32, 0, 0, 1.4)

	mtlMatte2 := rango.Material{}
	rango.SetMaterial(&mtlMatte2, blue, 0.1, 0.5, 0.4, 2, 0, 0, 1.4)

	mtlMatte3 := rango.Material{}
	rango.SetMaterial(&mtlMatte3, green, 0.1, 0.5, 0.4, 2, 0, 0, 1.4)

	mtlMatte4 := rango.Material{}
	rango.SetMaterial(&mtlMatte4, purple, 0.1, 0.5, 0.4, 2, 0, 0, 1.4)

	mtlGlass := rango.Material{}
	rango.SetMaterial(&mtlGlass, white, 0, 0.5, 0, 0, 0, 0.9, 1.4)

	/*Objects*/

	sphere := rango.Object{}
	rango.Sphere(&sphere, mtlMatte1, 0.9, 3)

	cone := rango.Object{}
	rango.Cone(&cone, mtlMatte4, 1, 1.5, 32)

	cylinder := rango.Object{}
	rango.Cylinder(&cylinder, mtlShiny2, 0.75, 1, 32)

	cube := rango.Object{}
	rango.Cube(&cube, mtlGlass, 1)

	planeBase := rango.Object{}
	rango.PlaneXZ(&planeBase, mtlShiny1, 10.0)

	planeLeft := rango.Object{}
	rango.PlaneXZ(&planeLeft, mtlMatte1, 10.0)

	planeRight := rango.Object{}
	rango.PlaneXZ(&planeRight, mtlMatte2, 10.0)

	planeBack := rango.Object{}
	rango.PlaneXZ(&planeBack, mtlMatte3, 10.0)

	rango.TransformObject(&sphere, rango.Translate(-0, 0.9, -2.7))
	rango.TransformObject(&cone, rango.Translate(2, 0, -2.8))
	rango.TransformObject(&cube, rango.MatrixMultiply(rango.Translate(2, 0.5, -1.0), rango.RotateY(45.0)))

	rango.TransformObject(&cylinder, rango.Translate(-0.0, 0.0, -1))
	rango.TransformObject(&planeBase, rango.Translate(1, 0, -4))
	rango.TransformObject(&planeLeft, rango.MatrixMultiply(rango.Translate(-2, 0, -4), rango.RotateZ(-90)))
	rango.TransformObject(&planeRight, rango.MatrixMultiply(rango.Translate(4, 0, -4), rango.RotateZ(90)))
	rango.TransformObject(&planeBack, rango.MatrixMultiply(rango.Translate(1, 0, -6), rango.RotateX(90)))

	rango.InitScene(scene, 8)
	rango.AddObjectsToScene(scene, sphere)
	rango.AddObjectsToScene(scene, cube)
	rango.AddObjectsToScene(scene, cone)
	rango.AddObjectsToScene(scene, cylinder)
	rango.AddObjectsToScene(scene, planeBase)
	rango.AddObjectsToScene(scene, planeRight)
	rango.AddObjectsToScene(scene, planeLeft)
	rango.AddObjectsToScene(scene, planeBack)

	return scene
}

func trace(ray rango.Ray, scene rango.Scene, light rango.Light, depth int) rango.Vector {

	outputColorVector := rango.Vector{0, 0, 0}
	hit := rango.IntersectScene(ray, scene)

	if hit.ObjectId >= 0 {
		diffuse := rango.Diffuse(hit, scene, light)
		ambient := rango.Ambient(hit, scene, light)
		specular := rango.Specular(hit, scene, light)
		outputColorVector = rango.Add(ambient, rango.Add(diffuse, specular))

		if depth > 0 {
			/* collect color captured by reflected ray */
			reflColor := trace(rango.ReflectRay(hit), scene, light, depth-1)
			krefl := scene.Objects[hit.ObjectId].Material.Reflectivity
			outputColorVector = rango.Add(outputColorVector, rango.FloatVecMult(krefl, reflColor))

			/* collect color captured by refracted ray */
			refrColor := trace(rango.RefractRay(hit, scene.Objects[hit.ObjectId].Material.Ir), scene, light, depth - 1)
			krefr := scene.Objects[hit.ObjectId].Material.Translucency
			outputColorVector = rango.Add(outputColorVector, rango.FloatVecMult(krefr, refrColor))
		}

		/* reduce color by the shadow factor of the light */
		return rango.FloatVecMult(1.0-rango.TraceShadow(hit, scene, light), outputColorVector)
	}

	return outputColorVector
}

func main() {

	width := 1280.0
	height := 960.0

	/* Setup Scene */
	scene := rango.Scene{}
	CreateScene(&scene)

	/* Setup lights */
	lColor := rango.Color{255, 255, 255}
	lPosition := rango.Vector{-1, 4, 4}
	light := rango.Light{lPosition, lColor, 0.3}

	/*Setup Camera*/
	camPosition := rango.Vector{1, 2, 4}
	lookAt := rango.Vector{1, 0, -6}
	camera := rango.Camera{}
	rango.SetCamera(&camera, camPosition, lookAt, 45.0, width, height)

	/*setup image*/
	outputImage := rango.Image{}
	rango.InitImage(&outputImage, uint32(width), uint32(height))

	var ray rango.Ray

	fmt.Println("Tracing started")
	for j := 0; j < int(height); j++ {
		for i := 0; i < int(width); i++ {
			ray = rango.GenerateRay(i, j, camera)
			outputColor := rango.Vector2Color(trace(ray, scene, light, 3))
			rango.SetPixel(&outputImage, uint32(i), uint32(j), outputColor)
		}
	}
	fmt.Println("Tracing done")

	fmt.Println("Writing image to disk")
	rango.WriteImage(&outputImage, "output.ppm")
	fmt.Println("Done writing the file to disk")

}
