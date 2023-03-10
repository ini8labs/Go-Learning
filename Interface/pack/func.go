package pack

import (
	"fmt"
	"math"
)

func (a Cube) Area() float32 {
	return 6 * a.side * a.side
}

func (a Cube) Volume() float32 {
	return a.side * a.side * a.side
}

func (b Cylinder) Area() float32 {
	return math.Pi * b.radius * b.radius * b.height
}

func (b Cylinder) Volume() float32 {
	return math.Pi * b.radius * b.radius * b.height
}

func (c Cuboid) Area() float32 {
	return 2 * (c.length*c.breadth + c.breadth*c.height + c.height*c.length)
}

func (c Cuboid) Volume() float32 {
	return c.length * c.breadth * c.height
}

func (d Sphere) Area() float32 {
	return 4 * math.Pi * d.side * d.side
}
func (d Sphere) Volume() float32 {
	return (4 / 3) * math.Pi * d.side * d.side * d.side
}

func calculate(o []Shape) {
	for _, v := range o {
		area := v.Area()
		volume := v.Volume()
		fmt.Println(v)
		fmt.Println("Area of object ", area)
		fmt.Println("Volume of object ", volume)
	}
}
