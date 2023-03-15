package shapes

import (
	"math"
)

type Sphere struct {
	Radius float32
}

func (d Sphere) Area() float32 {
	return float32(math.Round(float64(4 * math.Pi * d.Radius * d.Radius)))
}

func (d Sphere) Volume() float32 {
	return float32(math.Round(float64((4 / 3) * math.Pi * d.Radius * d.Radius * d.Radius)))
}
