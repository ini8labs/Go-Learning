package main

import "github.com/ini8labs/Go-Learning/Interface/InterfaceExample/shapes"

func main() {
	Cube := shapes.Cube{Side: 10}

	Cuboid := shapes.Cuboid{
		Length:  3,
		Breadth: 4,
		Height:  10,
	}

	Cylinder := shapes.Cylinder{
		Radius: 10,
		Height: 20,
	}

	Sphere := shapes.Sphere{
		Radius: 10,
	}

	object := []shapes.Shape{Cube, Cuboid, Cylinder, Sphere}
	shapes.Calculate(object)
}
