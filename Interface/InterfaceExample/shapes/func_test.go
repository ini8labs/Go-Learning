package shapes

import (
	"reflect"
	"testing"
)

func TestCalculate(t *testing.T) {
	type args struct {
		s []Shape
	}
	tests := []struct {
		name  string
		args  args
		want  []float32
		want1 []float32
	}{
		{
			name:  "Calcualte area and volume",
			args:  {cube, cuboid, cylinder, sphere},
			want:  {600, 190, 12566, 1018},
			want1: {1000, 150, 3770, 2290},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Calculate(tt.args.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Calculate() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Calculate() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
