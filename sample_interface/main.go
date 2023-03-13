package main

import (
	"Go-Learning/sample_interface/abc"
	"fmt"
)

func main() {

	var temp abc.Price
	temp = abc.Beemer{Name: "BMW", Yom: 2020, BasePrice: 300000}
	fmt.Printf("Current price of car : %s is %g ", temp.BrandName(), temp.PriceCalculator())

}
