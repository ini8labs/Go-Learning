package abc

import (
	"fmt"
	"time"
)

type Beemer struct {
	Name      string
	Yom       int
	BasePrice float64
}

func (a Beemer) PriceCalculator() float64 {
	life := time.Now().Year() - a.Yom
	depriciation := a.BasePrice * 0.03 * float64(life)
	FinalPrice := a.BasePrice - depriciation

	fmt.Println("Origional price of car : ", a.BasePrice)
	fmt.Println("Depriciation over time: ", depriciation)

	return FinalPrice
}

func (a Beemer) BrandName() string {
	return a.Name
}

type Toyota struct {
	Name      string
	Yom       int
	BasePrice float64
}

func (a Toyota) PriceCalculator() float64 {
	Life := time.Now().Year() - a.Yom
	depriciation := a.BasePrice * 0.1 * float64(Life)
	FinalPrice := a.BasePrice - depriciation
	fmt.Println("Origional price of car: ", a.BasePrice)
	fmt.Println("Depriciation over time: ", depriciation)
	return FinalPrice
}

func (a Toyota) BrandName() string {
	return a.Name
}
