package learnGoWithTests

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("saying hello to one", func(t *testing.T) {
		name := "zihong"
		got := Hello(name, "")
		want := "Hello, " + name

		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to empty", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Orange", "French")
		want := "Bonjour, Orange"
		assertCorrectMessage(t, got, want)
	})
}

func ExampleHello() {
	result := Hello("my", "")
	fmt.Println(result)
	// Output: Hello, my
}
