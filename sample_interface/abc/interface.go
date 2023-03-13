package abc

type Price interface {
	PriceCalculator() float64
	BrandName() string
}
