package shapes

import (
	"fmt"
)

var a = make([]float32, 4)
var vol = make([]float32, 4)

func Calculate(s []Shape) ([]float32, []float32, error) {

	for i, v := range s {
		area, err := v.Area()
		if err != nil {
			return nil, nil, err
		}
		volume, err := v.Volume()
		if err != nil {
			return nil, nil, err
		}

		a[i] = area
		vol[i] = volume

		fmt.Println("Dimensions of object: ", v)
		fmt.Println("Area of object: ", area)
		fmt.Println("Volume of object: ", volume)

	}
	return a, vol, nil
}
