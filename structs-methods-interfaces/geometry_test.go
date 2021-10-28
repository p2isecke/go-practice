package structs_methods_interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	// This is a function
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestAreaUsingTable(t *testing.T) {
	areaTests := []struct {
		name string
		shape Shape
		hasArea float64
	}{
		// run specific tests: `go test -run TestArea/Rectangle`
		{name: "Rectangle", shape :Rectangle{length:12, width: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{radius:10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{base:12, height: 6}, hasArea: 36},
	}

	// iterate over each test case
	for _, tc := range areaTests {
		// use t.Run and to name the test cases
		t.Run(tc.name, func(t *testing.T) {
			got := tc.shape.Area()

			if got != tc.hasArea {
				// %#v: prints out the struct with all its fields
				t.Errorf("%#v got %g, want %g", tc.shape, got, tc.hasArea)
			}
		})
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		// This is a method
		got := shape.Area()

		// %g allows us to print more precise decimals
		if got != want {
			t.Errorf("got %g, want %g", got, want)
		}
	}


	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		want := 72.0
		checkArea(t, rectangle, want)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		want := 314.1592653589793
		checkArea(t, circle, want)
	})
}