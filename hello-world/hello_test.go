package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to the people", func(t *testing.T) {
		spanishResultGot := Hello("Maxim", "Spanish")
		spanishResultWant := "Hola, Maxim"

		frenchResultGot := Hello("Maxim", "French")
		frenchResultWant := "Bonjour, Maxim"

		assertCorrectMessage(t, spanishResultGot, spanishResultWant)
		assertCorrectMessage(t, frenchResultGot, frenchResultWant)
	})

	t.Run("saying 'Hello, World!', if empty string is passed", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("if language not recognized, English is set by default", func(t *testing.T) {
		got := Hello("Maxim", "Umba-yumba")
		want := "Hello, Maxim"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
