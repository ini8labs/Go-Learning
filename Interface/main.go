package main

import (
	"fmt"

	"github.com/ini8labs/Go-Learning/Interface/carPrice"
)

func totalExpense(s []carPrice.Price) float64 {

	total := 0.0
	for _, v := range s {
		total += v.PriceCalculator()
	}
	return total
}

func main() {
	beemer1 := carPrice.Beemer{
		Name:      "BMW",
		Yom:       2020,
		BasePrice: 300000}

	toyota1 := carPrice.Toyota{
		Name:      "Corola",
		Yom:       2010,
		BasePrice: 400000}

	beemer2 := carPrice.Beemer{
		Name:      "M5 Competition",
		Yom:       2019,
		BasePrice: 100000}

	toyota2 := carPrice.Toyota{
		Name:      "Supra",
		Yom:       2000,
		BasePrice: 900000}

	inventory := []carPrice.Price{beemer1, beemer2, toyota1, toyota2}

	fmt.Println("The total asset price of the cars in inventory: ", totalExpense(inventory))

}
