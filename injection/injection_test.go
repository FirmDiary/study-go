package injection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Azal")

	got := buffer.String()
	want := "Hello, Azal"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}

}
