package main

import (
	"github.com/bnaveenkr/rango/rango"
)

func main()  {
	//vector := rango.Vector{1, 0, 0}
	//fmt.Println(vector)
	//var length float64 = rango.Length(vector)
	//print(length)

	color := rango.Color{255, 0, 0 }

	image := rango.Image{}
	newImage := rango.InitImage(&image, 100, 100)
	height := 100
	width := 100
	for j:=0; j<height; j++ {
		for i:=0; i<width; i++ {
			rango.SetPixel(newImage, uint32(i), uint32(j), color)
		}
	}

	rango.WriteImage(&image, "test.ppm")

}
