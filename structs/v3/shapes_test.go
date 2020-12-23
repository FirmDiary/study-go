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
	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Area()
		want := 100.0

		if got != want {
			t.Errorf("want %.2f but got %.2f", want, got)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.16

		if got != want {
			t.Errorf("got %.2f but want %2.f", got, want)
		}
	})

}
