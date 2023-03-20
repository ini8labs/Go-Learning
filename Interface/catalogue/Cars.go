package catalogue

type Cars interface {
	PriceCalculator() float64
	BrandName() string
}

func TotalExpense(s []Cars) float64 {

	total := 0.0
	for _, v := range s {
		total += v.PriceCalculator()
	}
	return total
}
