package v1

import "testing"

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"test": "this is just a test"}

	//t.Run("known word", func(t *testing.T) {
	//    got, _ := dictionary.Search("test")
	//    want := "this is just a test"
	//
	//    assertStrings(t, got, want)
	//})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		assertStrings(t, got, ErrNotFound)
	})

}

func assertStrings(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s' given, '%s'", got, want, "test")
	}

}
