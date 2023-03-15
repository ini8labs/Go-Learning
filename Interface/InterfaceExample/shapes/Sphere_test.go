package shapes

import (
	"testing"
)

func TestSphere_Area(t *testing.T) {
	tests := []struct {
		name string
		d    Sphere
		want float32
	}{
		{
			name: "Calculate area of Sphere",
			d: Sphere{
				Radius: 10,
			},
			want: 1257,
		},
		{
			name: "Checking area for -ve values",
			d: Sphere{
				Radius: -10,
			},
			want: 1257,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Area(); got != tt.want {
				t.Errorf("Sphere.Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSphere_Volume(t *testing.T) {
	tests := []struct {
		name string
		d    Sphere
		want float32
	}{
		{
			name: "Calculate volume of Sphere",
			d: Sphere{
				Radius: 10,
			},
			want: 3142,
		},
		{
			name: "Checking volume for -ve values",
			d: Sphere{
				Radius: -10,
			},
			want: -3142,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Volume(); got != tt.want {
				t.Errorf("Sphere.Volume() = %v, want %v", got, tt.want)
			}
		})
	}
}
