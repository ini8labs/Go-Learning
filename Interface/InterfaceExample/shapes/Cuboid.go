package shapes

import (
	"errors"
	"math"
)

type Cuboid struct {
	Length  float32
	Breadth float32
	Height  float32
}

func (c Cuboid) Area() (float32, error) {
	if c.Length < 0 || c.Breadth < 0 || c.Height < 0 {
		return -1, errors.New("can't calculate for -ve input")
	}
	return float32(math.Round(float64(2 * (c.Length*c.Breadth + c.Breadth*c.Height + c.Height*c.Length)))), nil
}

func (c Cuboid) Volume() (float32, error) {
	if c.Length < 0 || c.Breadth < 0 || c.Height < 0 {
		return -1, errors.New("can't calculate for -ve input")
	}
	return float32(math.Round(float64(c.Breadth * c.Length * c.Height))), nil
}
