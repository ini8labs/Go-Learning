package main

import (
	"Interface/pack"
)

func main() {
	Cube := pack.Cube{Side: 5}

	Cuboid := pack.Cuboid{
		Length:  5,
		Breadth: 4,
		Height:  10,
	}

	Cylinder := pack.Cylinder{
		Radius: 10,
		Height: 20,
	}

	Sphere := pack.Sphere{
		Radius: 10,
	}

	shapes := []pack.Shape{Cube, Cuboid, Cylinder, Sphere}
	pack.Calculate(shapes)

}
