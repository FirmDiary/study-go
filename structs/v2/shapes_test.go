package v1

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("want %.2f but got %.2f", want, got)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Area(rectangle)
	want := 100.0

	if got != want {
		t.Errorf("want %.2f but got %.2f", want, got)
	}
}
