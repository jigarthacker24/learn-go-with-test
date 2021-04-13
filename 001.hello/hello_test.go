package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, wnt string) {
		t.Helper()
		if got != wnt {
			t.Errorf("Output is not as expected. got: %q, want: %q", got, wnt)
		}
	}

	t.Run("Hello to people", func(t *testing.T) {
		got := Hello("Jigar", "")
		wnt := "Hello, Jigar"

		assertCorrectMessage(t, got, wnt)
	})

	t.Run("Hello with empty param", func(t *testing.T) {
		got := Hello("", "")
		wnt := "Hello, World"

		assertCorrectMessage(t, got, wnt)
	})

	t.Run("Hello in Spanish", func(t *testing.T) {
		got := Hello("Jigar", "Spanish")
		wnt := "Hola, Jigar"

		assertCorrectMessage(t, got, wnt)
	})

	t.Run("Hello in French", func(t *testing.T) {
		got := Hello("Jigar", "French")
		wnt := "Bonjour, Jigar"

		assertCorrectMessage(t, got, wnt)
	})
}
