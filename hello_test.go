package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Ran")
	want := "Hello Ran"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
