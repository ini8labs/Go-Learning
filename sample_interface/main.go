package main

import (
	"Go-Learning/sample_interface/abc"
	"fmt"
)

func main() {

	var temp abc.Price
	temp = abc.Beemer{Name: "BMW", Yom: 2020, Base_price: 300000}
	fmt.Printf("Current price of car : %s is %g ", temp.Brand_name(), temp.Price_calculator())

}
