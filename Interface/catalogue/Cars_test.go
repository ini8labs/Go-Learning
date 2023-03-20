package catalogue

import "testing"

func TestTotalExpense(t *testing.T) {
	// type args struct {
	// 	s []Cars
	// }

	bmw := catalogue.BWM {
		Name: "M5",
		Yom: 2017,
		BasePrice: 200000,
	}

	tests := []struct {
		name string
		args []catalogue.Cars
		want float64
	}
	
	{
		{name: "test1",
		args: []catalogue.Cars{bmw},
		want: 134000,
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TotalExpense(tt.args); got != tt.want {
				t.Errorf("TotalExpense() = %v, want %v", got, tt.want)
			}
		})
	}
}
