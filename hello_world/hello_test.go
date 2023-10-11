package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to daniel", func(t *testing.T) {
		expected := "Hello Daniel!"
		result := Hello("Daniel", "")
		assertExpectedResult(t, expected, result)
	})

	t.Run("saying hello world when empty string is provided", func(t *testing.T) {
		expected := "Hello world!"
		result := Hello("", "")
		assertExpectedResult(t, expected, result)
	})

	t.Run("say hello daniel in Spanish", func(t *testing.T) {
		expected := "Hola Daniel!"
		result := Hello("Daniel", "Spanish")
		assertExpectedResult(t, expected, result)
	})

	t.Run("say hello daniel in French", func(t *testing.T) {
		expected := "Bonjour Daniel!"
		result := Hello("Daniel", "French")
		assertExpectedResult(t, expected, result)
	})
}

func assertExpectedResult(t testing.TB, expected, result string) {
	t.Helper()

	if expected != result {
		t.Errorf("expected: %q result: %q", expected, result)
	}
}
