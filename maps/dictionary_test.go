package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		assertError(t, got, ErrNotFoundWord)
	})
}

func TestAdd(t *testing.T) {
	t.Run("add new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "test value"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("add new word", func(t *testing.T) {
		word := "test"
		definition := "test value"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "another test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test word"
		definition := "test definition"
		dictionary := Dictionary{word: definition}

		err := dictionary.Update(word, "new definition")

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, "new definition")
	})

	t.Run("non existed word", func(t *testing.T) {
		word := "test word"
		definition := "test definition"
		dictionary := Dictionary{word: definition}

		err := dictionary.Update("not existed word", "new definition")

		assertError(t, err, ErrWordNotExists)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)
	_, err := dictionary.Search(word)

	assertError(t, err, ErrNotFoundWord)
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal(err)
	}

	assertStrings(t, got, definition)
}
