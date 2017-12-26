package main

import (
	"fmt"
	"github.com/bnaveenkr/rango/rango"
)

func main()  {
	vector := rango.Vector{1, 0, 0}
	fmt.Println(vector)
	var length float64 = rango.Length(vector)
	print(length)
}
