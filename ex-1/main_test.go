package main

import "testing"

func TestSayHello(t *testing.T) {
	name := "Ben"
	want := GREETING_ENGLISH + name
	got := sayHello(LANGUAGE_ENGLISH, name)
	if want != got {
		t.Errorf("got %v want %v", got, want)
	}
}

// Sub tests
func TestSayHelloSubTests(t *testing.T) {
	// helper function for assertion
	assertHello := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("Test greeting in english", func(t *testing.T) {
		name := "Ben"
		want := GREETING_ENGLISH + name
		got := sayHello(LANGUAGE_ENGLISH, name)
		assertHello(t, got, want)
	})

	t.Run("Test greeting in french", func(t *testing.T) {
		name := "Beca"
		want := GREETING_FRENCH + name
		got := sayHello(LANGUAGE_FRENCH, name)
		assertHello(t, got, want)
	})

	t.Run("Test greeting in spanish", func(t *testing.T) {
		name := "Julie"
		want := GREETING_SPANISH + name
		got := sayHello(LANGUAGE_SPANISH, name)
		assertHello(t, got, want)
	})
}
