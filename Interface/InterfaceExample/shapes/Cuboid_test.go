package shapes

import (
	"testing"
)

func TestCuboid_Area(t *testing.T) {
	tests := []struct {
		name    string
		c       Cuboid
		want    float32
		wantErr bool
	}{
		{
			name: "Calculate Area of Cuboid",
			c: Cuboid{
				Length:  2,
				Breadth: 3,
				Height:  4,
			},
			want:    52,
			wantErr: false,
		},
		{
			name: "Checking Area for -ve values",
			c: Cuboid{
				Length:  2,
				Breadth: -3,
				Height:  4,
			},
			want:    -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Area()
			if (err != nil) != tt.wantErr {
				t.Errorf("Cuboid.Area() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Cuboid.Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCuboid_Volume(t *testing.T) {
	tests := []struct {
		name    string
		c       Cuboid
		want    float32
		wantErr bool
	}{
		{
			name: "Calculate Volume of Cuboid",
			c: Cuboid{
				Length:  2,
				Breadth: 3,
				Height:  4,
			},
			want:    24,
			wantErr: false,
		},
		{
			name: "Checking Volume for -ve values",
			c: Cuboid{
				Length:  2,
				Breadth: -3,
				Height:  4,
			},
			want:    -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Volume()
			if (err != nil) != tt.wantErr {
				t.Errorf("Cuboid.Volume() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Cuboid.Volume() = %v, want %v", got, tt.want)
			}
		})
	}
}
