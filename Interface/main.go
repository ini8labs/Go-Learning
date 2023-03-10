package main

import ("D:\office_git\Go-Learning\Interface\pack")

func main() {
	shape1 := Cube{5}
	shape2 := Cuboid{3, 4, 5}
	shape3 := Cylinder{3, 4}
	shape4 := Sphere{6}
	shapes := []Shape{shape1, shape2, shape3, shape4}
	calculate(shapes)

}
