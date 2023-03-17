package shapes

import (
	"errors"
	"math"
)

type Cylinder struct {
	Height float32
	Radius float32
}

func (b Cylinder) Area() (float32, error) {
	if b.Height < 0 || b.Radius < 0 {
		return -1, errors.New("can't calculate for -ve values")
	}
	return float32(math.Round(float64(math.Pi * b.Radius * b.Radius * b.Height))), nil
}

func (b Cylinder) Volume() (float32, error) {
	if b.Height < 0 || b.Radius < 0 {
		return -1, errors.New("can't calculate for -ve values")
	}
	return float32(math.Round(float64(2 * math.Pi * b.Radius * (b.Radius + b.Height)))), nil
}
