package shapes

import "testing"

func TestCube_Area(t *testing.T) {
	testCases := []struct {
		name   string
		inpt   Cube
		output float32
	}{
		{
			name: "Calculate Area of Cube",
			inpt: Cube{
				Side: 10,
			},
			output: 600,
		},

		{
			name: "Checking Area for -ve values",
			inpt: Cube{
				Side: -10,
			},
			output: 600,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.inpt.Area(); got != tt.output {
				t.Errorf("Cube.Area() = %v, want %v", got, tt.output)
			}
		})
	}
}

func TestCube_Volume(t *testing.T) {
	tests := []struct {
		name string
		inpt Cube
		otpt float32
	}{
		{
			name: "Calculate Volume of cube",
			inpt: Cube{
				Side: 10,
			},
			otpt: 1000,
		},
		{
			name: "Checking Volume for -ve values",
			inpt: Cube{
				Side: -10,
			},
			otpt: -1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.inpt.Volume(); got != tt.otpt {
				t.Errorf("Cube.Volume() = %v, want %v", got, tt.otpt)
			}
		})
	}
}
