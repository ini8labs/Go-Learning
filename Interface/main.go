package main

import (
	"fmt"

	"github.com/ini8labs/Go-Learning/Interface/catalogue"
)

func main() {
	beemer1 := catalogue.BWM{
		Name:      "BMW",
		Yom:       2020,
		BasePrice: 300000}

	toyota1 := catalogue.Toyota{
		Name:      "Corola",
		Yom:       2010,
		BasePrice: 400000}

	beemer2 := catalogue.BWM{
		Name:      "M5 Competition",
		Yom:       2019,
		BasePrice: 100000}

	toyota2 := catalogue.Toyota{
		Name:      "Supra",
		Yom:       2000,
		BasePrice: 900000}

	inventory := []catalogue.Cars{beemer1, beemer2, toyota1, toyota2}

	fmt.Println("The total asset price of the cars in inventory: ", catalogue.TotalExpense(inventory))

}
