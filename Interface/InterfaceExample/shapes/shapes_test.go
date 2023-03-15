package shapes

import (
	"testing"
)

func TestCube(t *testing.T) {

	shape := Cube{Side: 10}

	area := Cube.Area(shape)
	volume := Cube.Volume(shape)

	if area != 600 {
		t.Error("Incorrect area of cube: expected 600, got", area)
	}
	if volume != 1000 {
		t.Error("Incorrect area of cube: expected 1000, got", volume)
	}

}
func TestCuboid(t *testing.T) {

	shape := Cuboid{Length: 2, Breadth: 3, Height: 4}

	area := Cuboid.Area(shape)
	volume := Cuboid.Volume(shape)

	if area != 52 {
		t.Error("Incorrect area of cuboid: expected 52, got", area)
	}
	if volume != 24 {
		t.Error("Incorrect area of cuboid: expected 24, got", volume)
	}
}

func TestCylinder(t *testing.T) {

	shape := Cylinder{Height: 7, Radius: 6}

	area := Cylinder.Area(shape)
	volume := Cylinder.Volume(shape)

	if area != 792 {
		t.Error("Incorrect area of cuboid: expected 792, got", area)
	}
	if volume != 490 {
		t.Error("Incorrect area of cuboid: expected 490, got", volume)
	}
}

func TestSphere(t *testing.T) {

	shape := Sphere{Radius: 10}

	area := Sphere.Area(shape)
	volume := Sphere.Volume(shape)

	if area != 1257 {
		t.Error("Incorrect area of cuboid: expected 1257, got", area)
	}
	if volume != 3142 {
		t.Error("Incorrect area of cuboid: expected 3142, got", volume)
	}

}
