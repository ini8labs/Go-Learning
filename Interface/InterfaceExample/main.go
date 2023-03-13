package main

import "Go-Learning/Interface/Interface/InterfaceExample/pack"

func main() {
	Cube := pack.Cube{Side: 10}

	Cuboid := pack.Cuboid{
		Length:  3,
		Breadth: 5,
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
