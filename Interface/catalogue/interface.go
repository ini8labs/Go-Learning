package carPrice

type Price interface {
	PriceCalculator() float64
	BrandName() string
}
