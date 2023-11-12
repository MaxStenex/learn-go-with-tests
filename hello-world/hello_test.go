package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to the people", func(t *testing.T) {
		got := Hello("Maxim")
		want := "Hello, Maxim"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Hello, World!', if empty string is passed", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
