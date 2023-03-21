package catalogue

import "testing"

func TestBWM_PriceCalculator(t *testing.T) {
	tests := []struct {
		name string
		b    BWM
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.PriceCalculator(); got != tt.want {
				t.Errorf("BWM.PriceCalculator() = %v, want %v", got, tt.want)
			}
		})
	}
}
