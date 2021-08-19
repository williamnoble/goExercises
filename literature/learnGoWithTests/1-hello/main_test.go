package main

import "testing"

// Requires x_test.go, func must start with Test, one arg *testing.T

func TestHello(t *testing.T) {

	// Note: t is not a pointer
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)

		}
	}

	t.Run("Hello to a named individual", func(t *testing.T) {
		got := Hello("William", "")
		want := "Hello, William"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Hello World when a name parameter is not supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)

	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("William", "Spanish")
		want := "Hola, William"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("William", "French")
		want := "Bonjour, William"
		assertCorrectMessage(t, got, want)
	})

}
