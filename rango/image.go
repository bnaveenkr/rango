package rango

import (
	"fmt"
	"os"
)

const RGB uint8 = 3
const IMAGE_COMMENT = "Created using Rango"

type Image struct {
	Width  uint32
	Height uint32
	Data   []uint8
}

func InitImage(img *Image, width uint32, height uint32) *Image {
	img.Width = width
	img.Height = height

	pixels := (img.Width+1) * (img.Height+1) * 3

	img.Data = make([]uint8, pixels)
	return img
}

func SetPixel(img *Image, i uint32, j uint32, color Color) {

	h := img.Height
	w := img.Width
	var index uint32 = ((h-j)*w + i) * 3

	img.Data[index+0] = color.R
	img.Data[index+1] = color.G
	img.Data[index+2] = color.B
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func WriteImage(img *Image, filename string) {
	fp, err := os.Create(filename)
	check(err)

	defer fp.Close()
	headerString := fmt.Sprintf("P3\n#%s\n%d %d\n%d\n", IMAGE_COMMENT, img.Width, img.Height, 1<<8-1)
	fp.WriteString(headerString)

	width := int(img.Width)
	height := int(img.Height)

	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			index := (j*width + i) * int(RGB)
			pixelString := fmt.Sprintf("%d %d %d ", img.Data[index+0], img.Data[index+1], img.Data[index+2])
			fp.WriteString(pixelString)
		}
	}

	fp.Sync()
}
