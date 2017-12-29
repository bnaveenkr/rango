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

	cone := rango.Object{}
	rango.Cone(&cone, mtlMatte4, 1, 1.5, 32)

	cylinder := rango.Object{}
	rango.Cylinder(&cylinder, mtlShiny2, 0.75, 1, 32)

	cube := rango.Object{}
	rango.Cube(&cube, mtlGlass, 1)

	planeBase := rango.Object{}
	rango.PlaneXZ(&planeBase, mtlShiny1, 10)

	planeLeft := rango.Object{}
	rango.PlaneXZ(&planeLeft, mtlMatte1, 10)

	planeRight := rango.Object{}
	rango.PlaneXZ(&planeRight, mtlMatte2, 10)

	planeBack := rango.Object{}
	rango.PlaneXZ(&planeBack, mtlMatte3, 10)

	rango.TransformObject(&cone, rango.Translate(2, 0, -2.8))
	rango.TransformObject(&cube, rango.MatrixMultiply(rango.Translate(2, 0.5, -1.0), rango.RotateY(45.0)))
	rango.TransformObject(&cylinder, rango.Translate(0, 0, -1))
	rango.TransformObject(&planeBase, rango.Translate(1, 0, -4))
	rango.TransformObject(&planeLeft, rango.MatrixMultiply(rango.Translate(-2, 0, -4), rango.RotateZ(-90)))
	rango.TransformObject(&planeRight, rango.MatrixMultiply(rango.Translate(4, 0, -4), rango.RotateZ(90)))
	rango.TransformObject(&planeBack, rango.MatrixMultiply(rango.Translate(1, 0, -6), rango.RotateX(90)))

	rango.InitScene(scene, 8)
	rango.AddObjectsToScene(scene, cube)
	rango.AddObjectsToScene(scene, cone)
	rango.AddObjectsToScene(scene, cylinder)
	rango.AddObjectsToScene(scene, planeBack)
	rango.AddObjectsToScene(scene, planeRight)
	rango.AddObjectsToScene(scene, planeLeft)
	rango.AddObjectsToScene(scene, planeBack)

	return scene
}

func trace(ray rango.Ray, scene rango.Scene, light rango.Light) rango.Vector {

	outputColorVector := rango.Vector{0, 0, 0}
	hit := rango.IntersectScene(ray, scene)
	if hit.ObjectId >= 0 {
		diffuse := rango.Diffuse(hit, scene, light)
		ambient := rango.Ambinet(hit, scene, light)
		outputColorVector = rango.Add(diffuse, ambient)
	}
	fmt.Println(outputColorVector)
	return outputColorVector
}

func main() {

	width := 320.0
	height := 240.0

	/* Setup Scene */
	scene := rango.Scene{}
	CreateScene(&scene)

	/* Setup lights */
	lColor := rango.Color{255, 255, 255}
	lPosition := rango.Vector{01, 4, 4}
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
	for j := 0; j < int(height)-1; j++ {
		for i := 0; j < int(width)-1; i++ {
			ray = rango.GenerateRay(i, j, camera)
			rango.Vector2Color(trace(ray, scene, light))
			//fmt.Println(outputColor)
			//rango.SetPixel(&outputImage, uint32(i), uint32(j), outputColor)
		}
	}
	fmt.Println("Tracing done")

	fmt.Println("Writing image to disk")
	rango.WriteImage(&outputImage, "output.ppm")
	fmt.Println("Done writing the file to disk")

}
