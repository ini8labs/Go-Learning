package shapes

type Cuboid struct {
	Length  float32
	Breadth float32
	Height  float32
}

func (c Cuboid) Area() float32 {
	return 2 * (c.Length*c.Breadth + c.Breadth*c.Height + c.Height*c.Length)
}

func (c Cuboid) Volume() float32 {
	return c.Length * c.Breadth * c.Height
}
