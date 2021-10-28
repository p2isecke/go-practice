package main

import (
	"reflect"
	"testing"
)

func TestHello(t *testing.T) {
	// Pass in t *testing.T so that we can fail the test code
	// testing.TB is an interface that *testing.T and *testing.B uses
	assertCorrectMessage := func(t testing.TB, got, want string) {
		// Indicate to the test suite that this method is a helper
		// When there's a failure, it'll report the line number from the function call, not the test file
		t.Helper()

		if got != want {
			// Call Errorf METHOD on t to print out a message and fail the test
			// f allows us to format the string
			// %q are placeholders and wraps your values in double quotes
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("you", "")
		want := "Hello, you"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Maya", "Spanish")
		want := "Hola, Maya"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Jean", "French")
		want := "Bonjour, Jean"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Pig Latin", func(t *testing.T) {
		got := Hello("Dave", "Pig Latin")
		want := "Ellohay, Dave"

		assertCorrectMessage(t, got, want)
	})
}

// Using table driven tests: https://dave.cheney.net/2019/05/07/prefer-table-driven-tests
func TestHelloUsingTable(t *testing.T) {
	type test struct {
		name string
		got string
		want string
	}

	t.Helper()

	tests := []test{
		{name: "Empty String", got: Hello("", ""), want: "Hello, World"},
		{name: "With Name", got: Hello("you", ""), want: "Hello, you"},
		{name: "Spanish", got: Hello("Maya", "Spanish"), want: "Hola, Maya"},
		{name: "French", got: Hello("Jean", "French"), want: "Bonjour, Jean"},
		{name: "Pig Latin", got: Hello("Dave", "Pig Latin"), want: "Ellohay, Dave"},

	}

	for _, testCase := range tests {
		if !reflect.DeepEqual(testCase.want, testCase.got) {
			t.Errorf("got %q want %q", testCase.got, testCase.want)
		}
	}
}
