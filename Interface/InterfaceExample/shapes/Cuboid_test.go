package shapes

import (
	"testing"
)

func TestCuboid_Area(t *testing.T) {
	tests := []struct {
		name string
		c    Cuboid
		want float32
	}{
		{
			name: "Calculate Area of Cuboid",
			c: Cuboid{
				Length:  2,
				Breadth: 3,
				Height:  4,
			},
			want: 52,
		},
		{
			name: "Checking Area for -ve values",
			c: Cuboid{
				Length:  2,
				Breadth: -3,
				Height:  4,
			},
			want: -20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Area(); got != tt.want {
				t.Errorf("Cuboid.Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCuboid_Volume(t *testing.T) {
	tests := []struct {
		name string
		c    Cuboid
		want float32
	}{
		{
			name: "Calculate Volume of Cuboid",
			c: Cuboid{
				Length:  2,
				Breadth: 3,
				Height:  4,
			},
			want: 24,
		},
		{
			name: "Checking Volume for -ve values",
			c: Cuboid{
				Length:  2,
				Breadth: -3,
				Height:  4,
			},
			want: -24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Volume(); got != tt.want {
				t.Errorf("Cuboid.Volume() = %v, want %v", got, tt.want)
			}
		})
	}
}
