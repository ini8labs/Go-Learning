package shapes

import (
	"fmt"
)

func Calculate(s []Shape) {
	for _, v := range s {
		area := v.Area()
		volume := v.Volume()
		fmt.Println(v)
		fmt.Println("Area of object ", area)
		fmt.Println("Volume of object ", volume)
	}
}
