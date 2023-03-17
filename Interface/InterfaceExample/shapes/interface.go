package shapes

type Shape interface {
	Area() (float32, error)
	Volume() (float32, error)
}
