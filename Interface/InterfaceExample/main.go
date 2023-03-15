package main

import "github.com/ini8labs/Go-Learning/Interface/shapes"

func main() {
	cube := shapes.Cube{Side: 10}

	cuboid := shapes.Cuboid{
		Length:  3,
		Breadth: 5,
		Height:  10,
	}

	cylinder := shapes.Cylinder{
		Radius: 10,
		Height: 20,
	}

	sphere := shapes.Sphere{
		Radius: 9,
	}

	object := []shapes.Shape{cube, cuboid, cylinder, sphere}
	shapes.Calculate(object)
}
