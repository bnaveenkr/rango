package main

import (
	"fmt"
	"github.com/bnaveenkr/rango/rango"
)

func main() {
	//vector := rango.Vector{1, 0, 0}
	//fmt.Println(vector)
	//var length float64 = rango.Length(vector)
	//print(length)

	color := rango.Color{255, 0, 0}
	//
	//image := rango.Image{}
	//newImage := rango.InitImage(&image, 100, 100)
	//height := 100
	//width := 100
	//for j:=0; j<height; j++ {
	//	for i:=0; i<width; i++ {
	//		rango.SetPixel(newImage, uint32(i), uint32(j), color)
	//	}
	//}
	//
	//rango.WriteImage(&image, "test.ppm")

	cube := rango.Object{}
	material := rango.Material{}
	rango.SetMaterial(&material, color, 0.1, 0.5, 0.4, 32, 0, 0, 1.4)
	rango.Cube(&cube, material, 1)

	cone := rango.Object{}
	rango.Cone(&cone, material, 1, 1.5, 32)

	plane := rango.Object{}
	rango.PlaneXZ(&plane, material, 10)
	fmt.Println(plane)


	cylinder := rango.Object{}
	rango.Cylinder(&cylinder, material, 0.75, 1, 4)
	fmt.Println(cylinder)

}
