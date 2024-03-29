package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, actual, expected string) {
		t.Helper()

		if actual != expected {
			t.Errorf("Got %q, wanted %q", actual, expected)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Chris", "English")
		wanted := "Hello, Chris"

		assertCorrectMessage(t, got, wanted)
	})

	t.Run("greet stangers with World", func(t *testing.T) {
		got := Hello("", "")
		wanted := "Hello, World"

		assertCorrectMessage(t, got, wanted)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Jozsi", "Spanish")
		wanted := "Elodie, Jozsi"

		assertCorrectMessage(t, got, wanted)

	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Joska", "French")
		wanted := "Bonjour, Joska"
		assertCorrectMessage(t, got, wanted)
	})

}
