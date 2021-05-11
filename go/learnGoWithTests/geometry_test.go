package learnGoWithTests

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	areaTests := []struct{
		g Geometry
		want float64
	}{
		{&Rectangle{10,10.5}, 41},
		{&Circle{10}, 23.561944901923447},
		{&Triangle{6, 4}, 16},
	}
	for _, test := range areaTests {
		got := test.g.Perimeter()
		if got != test.want {
			t.Errorf("got %.2f but want %.2f", got, test.want)
		}
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct{
		g Geometry
		want float64
	}{
		{&Rectangle{10,10.5}, 105},
		{&Circle{10}, 314.1592653589793},
		{&Triangle{12, 6}, 36.0},
	}
	for _, test := range areaTests {
		got := test.g.Area()
		if got != test.want {
			t.Errorf("got %.2f but want %.2f", got, test.want)
		}
	}
}