package main


func main() {
	Cube := objects.Cube{Side: 10}

	Cuboid := objects.Cuboid{
		Length:  3,
		Breadth: 5,
		Height:  10,
	}

	Cylinder := objects.Cylinder{
		Radius: 10,
		Height: 20,
	}

	Sphere := objects.Sphere{
		Radius: 10,
	}

	shapes := []objects.Shape{Cube, Cuboid, Cylinder, Sphere}
	objects.Calculate(shapes)
}
