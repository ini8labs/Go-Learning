package shapes

import (
	"errors"
	"math"
)

type Cube struct {
	Side float32
}

func (a Cube) Area() (float32, error) {

	if a.Side < 0 {
		return -1, errors.New("can't calculate for -ve input")
	}

	return float32(math.Round(float64(6 * a.Side * a.Side))), nil

}

func (a Cube) Volume() (float32, error) {

	if a.Side < 0 {
		return -1, errors.New("can't calculate for -ve input")
	}

	return float32(math.Round(float64(a.Side * a.Side * a.Side))), nil

}
