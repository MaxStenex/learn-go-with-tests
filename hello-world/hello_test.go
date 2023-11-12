package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Maxim")
	want := "Hello, Maxim"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
