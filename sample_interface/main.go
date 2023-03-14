package main

import (
	"Go-Learning/sample_interface/abc"
	"fmt"
)

func totalExpense(s []abc.Price) float64 {

	Total := 0.0
	for _, v := range s {
		Total += v.PriceCalculator()
	}
	return Total
}

func main() {
	temp1 := abc.Beemer{
		Name:      "BMW",
		Yom:       2020,
		BasePrice: 300000}
	temp2 := abc.Toyota{
		Name:      "Corola",
		Yom:       2010,
		BasePrice: 400000}
	temp3 := abc.Beemer{
		Name:      "M5 Competition",
		Yom:       2019,
		BasePrice: 100000}
	temp4 := abc.Toyota{
		Name:      "Supra",
		Yom:       2000,
		BasePrice: 900000}

	inventory := []abc.Price{temp1, temp2, temp3, temp4}

	fmt.Println("The total price of the cars are: ", totalExpense(inventory))

}
