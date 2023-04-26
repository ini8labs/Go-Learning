package shapes

import (
	"testing"
)

func TestCylinder_Area(t *testing.T) {
	tests := []struct {
		name    string
		b       Cylinder
		want    float32
		wantErr bool
	}{
		{
			name: "Calculate area of Cylinder",
			b: Cylinder{
				Height: 7,
				Radius: 6,
			},
			want:    792,
			wantErr: false,
		},
		{
			name: "Checking area for -ve values",
			b: Cylinder{
				Height: -7,
				Radius: 6,
			},
			want:    -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.b.Area()
			if (err != nil) != tt.wantErr {
				t.Errorf("Cylinder.Area() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Cylinder.Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCylinder_Volume(t *testing.T) {
	tests := []struct {
		name    string
		b       Cylinder
		want    float32
		wantErr bool
	}{
		{
			name: "Calculate Volume of Cylinder",
			b: Cylinder{
				Height: 7,
				Radius: 6,
			},
			want:    490,
			wantErr: false,
		},
		{
			name: "Checking Volume for -ve values",
			b: Cylinder{
				Height: -7,
				Radius: 6,
			},
			want:    -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.b.Volume()
			if (err != nil) != tt.wantErr {
				t.Errorf("Cylinder.Volume() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Cylinder.Volume() = %v, want %v", got, tt.want)
			}
		})
	}
}
