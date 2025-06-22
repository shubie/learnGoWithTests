package main

import "testing"

func TestHelloWorld(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		got := HelloWorld("Shuaib", "english")
		want := "Hello, Shuaib"

		assertCorrectMessages(t, got, want)
	})

	t.Run("when the string is empty, display hello world", func(t *testing.T) {
		got := HelloWorld("", "english")
		want := "Hello, World"

		assertCorrectMessages(t, got, want)
	})

	t.Run("Display in spanish", func(t *testing.T) {
		got := HelloWorld("shuaib", "spanish")
		want := "Hola, shuaib"

		assertCorrectMessages(t, got, want)
	})

	t.Run("Display in french", func(t *testing.T) {
		got := HelloWorld("shuaib", "french")
		want := "Boujor, shuaib"

		assertCorrectMessages(t, got, want)
	})

}

func assertCorrectMessages(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
