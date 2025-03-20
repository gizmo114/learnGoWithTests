package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to people that provide their name", func(t *testing.T) {
		got := Hello("Tomas")
		want := "Hello, Tomas!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World!' when an empty string is provided", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world!"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// makes this a helper function so the error will be reported in the correct line
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
