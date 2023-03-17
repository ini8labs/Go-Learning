package main

import (
	"testing"

	"github.com/ini8labs/Go-Learning/Interface/catalogue"
)

func Test_totalExpense(t *testing.T) {

	bmw := catalogue.Beemer{
		Name:      "BMW",
		Yom:       2012,
		BasePrice: 10,
	}
	tests := []struct {
		name  string
		input []catalogue.Price
		want  float64
	}{
		{
			name:  "test",
			input: []catalogue.Price{bmw},
			want:  6.7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalExpense(tt.input); got != tt.want {
				t.Errorf("totalExpense() = %v, want %v", got, tt.want)
			}
		})
	}
}
