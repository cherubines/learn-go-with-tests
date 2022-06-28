package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("in English", func(t *testing.T) {
		got := Hello("Kyo", "English")
		want := "Hello, Kyo"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Kyo", "French")
		want := "Bonjour, Kyo"
		assertCorrectMessage(t, got, want)
	})
}
