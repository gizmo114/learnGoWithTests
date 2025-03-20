package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Tomas")
	want := "Hello, Tomas!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
