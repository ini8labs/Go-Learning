package main

import (
	"testing"

	"github.com/ini8labs/Go-Learning/Interface/catalogue"
)

func Test_totalExpense(t *testing.T) {
	type args struct {
		s []catalogue.Price
	}
	tests := []struct {
		name string
		inpt args
		outpt float64
	},{
		name: "beemer1",
		inpt: args.s.PriceCalculator{Name:"BMW",
		Yom:2020,
		BasePrice:300000}, 
		outpt: 273000
	 },
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if inpr := totalExpense(tt.args.s); inpt != tt.outpt {
				t.Errorf("totalExpense() = %v, want %v", inpt, tt.outpt)
			}
		})
	}
}
