package shapes

import (
	"fmt"
)

var a = make([]float32, 4)
var vol = make([]float32, 4)

func Calculate(s []Shape) ([]float32, []float32) {

	var area float32 = 0
	var volume float32 = 0
	for _, v := range s {
		area = v.Area()
		volume = v.Volume()
		a = append(a, area)
		vol = append(vol, volume)
		fmt.Println("Dimensions of object: ", v)
		fmt.Println("Area of object: ", area)
		fmt.Println("Volume of object: ", volume)

	}
	return a, vol
}
