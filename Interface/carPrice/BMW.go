package carPrice

import (
	"fmt"
	"time"
)

type Beemer struct {
	Name      string
	Yom       int
	BasePrice float64
}

func (b Beemer) PriceCalculator() float64 {
	life := time.Now().Year() - b.Yom
	depriciation := b.BasePrice * 0.03 * float64(life)
	FinalPrice := b.BasePrice - depriciation

	fmt.Println("Origional price of car : ", b.BasePrice)
	fmt.Println("Depriciation over time: ", depriciation)

	return FinalPrice
}

func (b Beemer) BrandName() string {
	return b.Name
}
