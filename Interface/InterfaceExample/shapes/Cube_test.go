package shapes

import (
	"testing"
)

func TestCube_Area(t *testing.T) {
	tests := []struct {
		name    string
		a       Cube
		want    float32
		wantErr bool
	}{
		{
			name: "Calculate Area of Cube",
			a: Cube{
				Side: 10,
			},
			want:    600,
			wantErr: false,
		},
		{
			name: "Checking Area for -ve values",
			a: Cube{
				Side: -10,
			},
			want:    -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.Area()
			if (err != nil) != tt.wantErr {
				t.Errorf("Cube.Area() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Cube.Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCube_Volume(t *testing.T) {
	tests := []struct {
		name    string
		a       Cube
		want    float32
		wantErr bool
	}{
		{
			name: "Calculate Volume of cube",
			a: Cube{
				Side: 10,
			},
			want:    1000,
			wantErr: false,
		},
		{
			name: "Checking Volume for -ve values",
			a: Cube{
				Side: -10,
			},
			want:    -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.Volume()
			if (err != nil) != tt.wantErr {
				t.Errorf("Cube.Volume() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Cube.Volume() = %v, want %v", got, tt.want)
			}
		})
	}
}
