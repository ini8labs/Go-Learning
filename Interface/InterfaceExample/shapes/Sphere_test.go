package shapes

import (
	"testing"
)

func TestSphere_Area(t *testing.T) {
	tests := []struct {
		name    string
		d       Sphere
		want    float32
		wantErr bool
	}{
		{
			name: "Calculate area of Sphere",
			d: Sphere{
				Radius: 10,
			},
			want:    1257,
			wantErr: false,
		},
		{
			name: "Checking area for -ve values",
			d: Sphere{
				Radius: -10,
			},
			want:    -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.Area()
			if (err != nil) != tt.wantErr {
				t.Errorf("Sphere.Area() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Sphere.Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSphere_Volume(t *testing.T) {
	tests := []struct {
		name    string
		d       Sphere
		want    float32
		wantErr bool
	}{
		{
			name: "Calculate volume of Sphere",
			d: Sphere{
				Radius: 10,
			},
			want:    3142,
			wantErr: false,
		},
		{
			name: "Checking volume for -ve values",
			d: Sphere{
				Radius: -10,
			},
			want:    -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.Volume()
			if (err != nil) != tt.wantErr {
				t.Errorf("Sphere.Volume() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Sphere.Volume() = %v, want %v", got, tt.want)
			}
		})
	}
}
