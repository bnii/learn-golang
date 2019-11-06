package structs

import "testing"

func assert(got, expected float64, t *testing.T) {
	t.Helper()
	if got != expected {
		t.Errorf("got %.2f expected %.2f", got, expected)
	}
}

func TestPerimiter(t *testing.T) {
	assert(Perimiter(Rectangle{3.0, 4.0}), 14.0, t)
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{12, 6}, want: 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{
			Height: 6,
			Base:   12,
		}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}

}
