package pack

import (
	"fmt"
	"math"
)

func (a Cube) Area() float32 {
	return 6 * a.Side * a.Side
}

func (a Cube) Volume() float32 {
	return a.Side * a.Side * a.Side
}

func (b Cylinder) Area() float32 {
	return math.Pi * b.Radius * b.Radius * b.Height
}

func (b Cylinder) Volume() float32 {
	return math.Pi * b.Radius * b.Radius * b.Height
}

func (c Cuboid) Area() float32 {
	return 2 * (c.Length*c.Breadth + c.Breadth*c.Height + c.Height*c.Length)
}

func (c Cuboid) Volume() float32 {
	return c.Length * c.Breadth * c.Height
}

func (d Sphere) Area() float32 {
	return 4 * math.Pi * d.Radius * d.Radius
}
func (d Sphere) Volume() float32 {
	return (4 / 3) * math.Pi * d.Radius * d.Radius * d.Radius
}

func Calculate(o []Shape) {
	for _, v := range o {
		area := v.Area()
		volume := v.Volume()
		fmt.Println(v)
		fmt.Println("Area of object ", area)
		fmt.Println("Volume of object ", volume)
	}
}
