package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("empty strings returns 'Hello, world'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people in English", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world' in English", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})
	
	t.Run("saying hello to people in Spanish", func(t *testing.T) {
		got := Hello("Pablo", "Spanish")
		want := "Hola, Pablo"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world' in Spanish", func(t *testing.T) {
		got := Hello("", "Spanish")
		want := "Hola, world"
		assertCorrectMessage(t, got, want)
	})
	
	t.Run("saying hello to people in French", func(t *testing.T) {
		got := Hello("Jean", "French")
		want := "Bonjour, Jean"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world' in French", func(t *testing.T) {
		got := Hello("", "French")
		want := "Bonjour, world"
		assertCorrectMessage(t, got, want)
	})
	
	t.Run("saying hello to people in Arabic", func(t *testing.T) {
		got := Hello("Abbas", "Arabic")
		want := "Marhaba, Abbas"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world' in Arabic", func(t *testing.T) {
		got := Hello("", "Arabic")
		want := "Marhaba, world"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}