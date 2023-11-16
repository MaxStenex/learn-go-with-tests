package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{12.0, 6.0}, want: 72},
		{shape: Circle{15}, want: 706.8583470577034},
		{shape: Triangle{12, 5}, want: 30},
	}

	for _, test := range areaTests {
		got := test.shape.Area()
		want := test.want

		if got != want {
			t.Errorf("%#v got %g want %g", test, got, want)
		}
	}
}
