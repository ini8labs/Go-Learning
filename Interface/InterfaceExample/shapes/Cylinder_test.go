package shapes

import (
	"testing"
)

func TestCylinder_Area(t *testing.T) {
	tests := []struct {
		name string
		b    Cylinder
		want float32
	}{
		{
			name: "Calculate area of Cylinder",
			b: Cylinder{
				Height: 7,
				Radius: 6,
			},
			want: 792,
		},
		{
			name: "Checking area for -ve values",
			b: Cylinder{
				Height: -7,
				Radius: 6,
			},
			want: -792,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Area(); got != tt.want {
				t.Errorf("Cylinder.Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCylinder_Volume(t *testing.T) {
	tests := []struct {
		name string
		b    Cylinder
		want float32
	}{
		{
			name: "Calculate Volume of Cylinder",
			b: Cylinder{
				Height: 7,
				Radius: 6,
			},
			want: 490,
		},
		{
			name: "Checking Volume for -ve values",
			b: Cylinder{
				Height: -7,
				Radius: 6,
			},
			want: -38,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Volume(); got != tt.want {
				t.Errorf("Cylinder.Volume() = %v, want %v", got, tt.want)
			}
		})
	}
}
