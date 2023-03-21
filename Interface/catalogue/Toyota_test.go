package catalogue

import "testing"

func TestToyota_PriceCalculator(t *testing.T) {

	toyota1 := catalogue.Toyota {
		Name: "fortuner", 
		Yom: 2016, 
		BasePrice: 50000,
		}

	tests := []struct {
		name string
		tr   Toyota
		want float64
	}

	{
		{
		name: "test1",
		tr: toyota1,
		want: 47000,
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PriceCalculator(tt.toyota1); got != tt.want {
				t.Errorf("Toyota.PriceCalculator() = %v, want %v", got, tt.want)
			}
		})
	}
}
