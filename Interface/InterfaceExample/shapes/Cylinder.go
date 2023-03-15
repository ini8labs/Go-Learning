package shapes

import (
	"math"
)

type Cylinder struct {
	Height float32
	Radius float32
}

func (b Cylinder) Area() float32 {
	return float32(math.Round(float64(math.Pi * b.Radius * b.Radius * b.Height)))
}

func (b Cylinder) Volume() float32 {
	return float32(math.Round(float64(2 * math.Pi * b.Radius * (b.Radius + b.Height))))
}
