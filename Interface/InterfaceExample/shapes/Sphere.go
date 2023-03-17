package shapes

import (
	"errors"
	"math"
)

type Sphere struct {
	Radius float32
}

func (d Sphere) Area() (float32, error) {
	if d.Radius <= 0 {
		return -1, errors.New("can't calculate for zero and-ve values")
	}
	return float32(math.Round(float64(4 * math.Pi * d.Radius * d.Radius))), nil
}

func (d Sphere) Volume() (float32, error) {
	if d.Radius <= 0 {
		return -1, errors.New("can't calculate for zero and -ve values")
	}
	return float32(math.Round(float64((4 / 3) * math.Pi * d.Radius * d.Radius * d.Radius))), nil
}
