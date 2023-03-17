package shapes

import (
	"reflect"
	"testing"
)

func TestCalculate(t *testing.T) {
	type calculateReturn struct {
		area   []float32
		volume []float32
		err    error
	}

	type calculateInput struct {
		s []Shape
	}

	cube := Cube{Side: 10}
	cuboid := Cuboid{Length: 3, Breadth: 5, Height: 10}
	cylinder := Cylinder{Height: 10, Radius: 20}
	sphere := Sphere{Radius: 9}

	input := calculateInput{
		s: []Shape{cube, cuboid, cylinder, sphere},
	}

	output := calculateReturn{
		area:   []float32{600, 190, 12566, 1018},
		volume: []float32{1000, 150, 3770, 2290},
		err:    nil,
	}

	tests := []struct {
		name    string
		args    calculateInput
		want    []float32
		want1   []float32
		wantErr bool
	}{
		{
			name:    "Calculate",
			args:    input,
			want:    output.area,
			want1:   output.volume,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := Calculate(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Calculate() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Calculate() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
