// Package shapes provides various utilities to calculate diffrent
// geometrical properties
package shapes

// Cube represents the measure of Side of the Cube.
type Cube struct {
	Side float32
}

// Area() is a method which calculates the area of Cube.
func (a Cube) Area() float32 {
	return 6 * a.Side * a.Side
}

// Volume() is a method which calculates the volume of Cube.
func (a Cube) Volume() float32 {
	return a.Side * a.Side * a.Side
}
