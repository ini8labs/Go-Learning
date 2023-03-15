package shapes

type Cube struct {
	Side float32
}

func (a Cube) Area() float32 {
	return 6 * a.Side * a.Side
}

func (a Cube) Volume() float32 {
	return a.Side * a.Side * a.Side
}
