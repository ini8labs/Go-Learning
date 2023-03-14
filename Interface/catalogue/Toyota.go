package carPrice

import (
	"fmt"
	"time"
)

type Toyota struct {
	Name      string
	Yom       int
	BasePrice float64
}

func (t Toyota) PriceCalculator() float64 {
	life := time.Now().Year() - t.Yom
	depriciation := t.BasePrice * 0.05 * float64(life)
	FinalPrice := t.BasePrice - depriciation

	fmt.Println("Origional price of car: ", t.BasePrice)
	fmt.Println("Depriciation over time: ", depriciation)

	return FinalPrice
}

func (t Toyota) BrandName() string {
	return t.Name
}
