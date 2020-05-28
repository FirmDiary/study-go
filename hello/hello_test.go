package main

import "testing"

func TestHello(t *testing.T){
    got := Hello("Azal")
    want := "Hello, Azal"

    if got != want {
        t.Errorf("got '%q' want '%q'", got, want)
    }
}
