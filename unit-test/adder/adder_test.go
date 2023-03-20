package adder

import (
	"testing"
)

func Test_addNumners(t *testing.T) {

	result := addNumebrs(1, 3)
	if result != 4 {
		t.Error("incorrect error: expected 5, got", result)
	}
}

func Test_addNumebrs(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addNumebrs(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("addNumebrs() = %v, want %v", got, tt.want)
			}
		})
	}
}
