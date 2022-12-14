package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {

	t.Run("rectangle", func(t *testing.T) {

		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Perimeter()
		want := 40.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("circle", func(t *testing.T) {

		circle := Circle{10.0}
		got := circle.Perimeter()
		want := 62.83185307179586

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})

}

func TestArea(t *testing.T) {

	t.Run("rectangle", func(t *testing.T) {

		rectangle := Rectangle{12.0, 6.0}

		got := rectangle.Area()
		want := 72.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}

		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()

		got := shape.Area()

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangle with checkArea", func(t *testing.T) {
		rectangle := Rectangle{12, 6}

		checkArea(t, rectangle, 72.0)
	})

	/**
	 *  Table driven tests
	 *  https://github.com/golang/go/wiki/TableDrivenTests
	 */

	t.Run("using table driven tests", func(t *testing.T) {

		areaTests := []struct {
			shape Shape
			want  float64
		}{
			{Rectangle{12, 6}, 72.0},
			{Circle{10.0}, 314.1592653589793},
			{Triangle{12, 6}, 36.0},
		}

		for _, tt := range areaTests {
			got := tt.shape.Area()

			if got != tt.want {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
			}
		}

	})
}
