package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{1, 6}
	p := Perimeter(rect)
	wnt := 14.0

	if p != wnt {
		t.Errorf("Perimeter not correct. got: %.2f, want:%.2f", p, wnt)
	}
}

func TestArea(t *testing.T) {

	_ = func(t testing.TB, shape Shape, wnt float64) {
		t.Helper()
		got := shape.Area()
		if got != wnt {
			t.Errorf("Area not correct. got: %g, want:%g", got, wnt)
		}
	}

	tests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 5, Height: 6}, hasArea: 30.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("Area not correct.shape: %#v, got: %g, want: %g", tt.shape, got, tt.hasArea)
			}

		})
	}

}
