package shapes

import "math"

type Cube struct {
	Side float32
}

func (a Cube) Area() float32 {
	return float32(math.Round(float64(6 * a.Side * a.Side)))
}

func (a Cube) Volume() float32 {
	return float32(math.Round(float64(a.Side * a.Side * a.Side)))
}
