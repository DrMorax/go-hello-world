package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("empty strings returns 'Hello, Traveler'",
	 func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, Traveler"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people in English",
	 func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world' in English",
	 func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, Traveler"
		assertCorrectMessage(t, got, want)
	})
	
	t.Run("saying hello to people in Spanish",
	 func(t *testing.T) {
		got := Hello("Pablo", "Spanish")
		want := "Hola, Pablo"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world' in Spanish",
	 func(t *testing.T) {
		got := Hello("", "Spanish")
		want := "Hola, Traveler"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}