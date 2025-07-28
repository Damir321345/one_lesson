package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Здороваться с людьми", func(t *testing.T) {
		got := Hello("Ran")
		want := "Hello Ran"
		assertCorrectMessage(t, got, want)
	})
	t.Run("С пустой строки", func(t *testing.T) {
		got := Hello("")
		want := "Hello world"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
