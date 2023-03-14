package shapes

import (
	"math"
)

type Cylinder struct {
	Height float32
	Radius float32
}

func (b Cylinder) Area() float32 {
	return math.Pi * b.Radius * b.Radius * b.Height
}

func (b Cylinder) Volume() float32 {
	return math.Pi * b.Radius * b.Radius * b.Height
}
