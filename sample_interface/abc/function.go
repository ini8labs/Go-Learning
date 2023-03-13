package abc

import (
	"fmt"
	"time"
)

type Beemer struct {
	Name       string
	Yom        int
	Base_price float64
}

func (a Beemer) Price_calculator() float64 {
	Life := time.Now().Year() - a.Yom
	Depriciation := a.Base_price * 0.03 * float64(Life)
	Final_price := a.Base_price - Depriciation
	fmt.Println("Origional price of car : ", a.Base_price)
	fmt.Println("Depriciation over time: ", Depriciation)
	return Final_price
}

func (a Beemer) Brand_name() string {
	return a.Name
}

type Toyota struct {
	Name       string
	Yom        int
	Base_price float64
}

func (a Toyota) Price_calculator() float64 {
	Life := time.Now().Year() - a.Yom
	Depriciation := a.Base_price * 0.1 * float64(Life)
	Final_price := a.Base_price - Depriciation
	fmt.Println("Origional price of car : ", a.Base_price)
	fmt.Println("Depriciation over time: ", Depriciation)
	return Final_price
}

func (a Toyota) Brand_name() string {
	return a.Name
}
