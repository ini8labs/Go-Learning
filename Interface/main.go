package main

import (
	"fmt"

	"github.com/ini8labs/Go-Learning/Interface/carPrice"
)

func totalExpense(s []carPrice.Price) float64 {

	Total := 0.0
	for _, v := range s {
		Total += v.PriceCalculator()
	}
	return Total
}

func main() {
	car1 := carPrice.Beemer{
		Name:      "BMW",
		Yom:       2020,
		BasePrice: 300000}
	car2 := carPrice.Toyota{
		Name:      "Corola",
		Yom:       2010,
		BasePrice: 400000}
	car3 := carPrice.Beemer{
		Name:      "M5 Competition",
		Yom:       2019,
		BasePrice: 100000}
	car4 := carPrice.Toyota{
		Name:      "Supra",
		Yom:       2000,
		BasePrice: 900000}

	Inventory := []carPrice.Price{car1, car2, car3, car4}

	fmt.Println("The total asset price of the cars in inventory: ", totalExpense(Inventory))

}
