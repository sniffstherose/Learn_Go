package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("Saying hello to peple", func(t *testing.T) {
		got := Hello("Santonlee", "")
		want := "Hello, Santonlee"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello world when an empty stirng is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("In Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In French", func(t *testing.T) {
		got := Hello("Salyd", "French")
		want := "Bonjour, Salyd"
		assertCorrectMessage(t, got, want)
	})
}
