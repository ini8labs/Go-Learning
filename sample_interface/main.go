package main

import (
	"Go-Learning/sample_interface/packages"
	"fmt"
)

func main() {

	var temp packages.Price
	temp = packages.Beemer{"BMW", 2020, 300000}
	fmt.Printf("Current price of car : %s is %g ", temp.Brand_name(), temp.Price_calculator())

}
