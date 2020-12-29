package v1

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		assertStrings(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {

	t.Run("new key", func(t *testing.T) {
		dictionary := Dictionary{}

		key := "test"
		word := "this is just a test"

		dictionary.Add(key, word)

		assertDefinition(t, dictionary, key, word)
	})

	t.Run("exist key", func(t *testing.T) {
		key := "test"
		word := "this is just a test"

		dictionary := Dictionary{key: word}

		err := dictionary.Add(key, "new test")

		assertError(t, err, ErrKeyExists)
		assertDefinition(t, dictionary, key, word)
	})

}

func assertError(t *testing.T, got error, want error) {
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if got != want {
		t.Errorf("got '%s', but want '%s'", got, want)
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, key, word string) {
	t.Helper()

	got, err := dictionary.Search(key)

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if word != got {
		t.Errorf("got '%s' but want '%s'", got, word)
	}
}

func assertStrings(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s' given, '%s'", got, want, "test")
	}

}
