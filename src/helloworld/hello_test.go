package helloworld

import "testing"

func TestHello(t *testing.T) {

	t.Run("saying Hello to people", func(t *testing.T) {
		got := Hello("Anna.","EN")
		want := "Hello, Anna."

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {

		got := Hello("","EN")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "ES")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})


	t.Run("in French", func(t *testing.T) {
		got := Hello("Anna", "FR")
		want := "Bonjour, Anna"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
