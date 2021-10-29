package dependency_injection

import (
	"bytes"
	"testing"
)

// we don't care where or how the printing happens so we can accept an interface instead of a concrete type
// when you call fmt.Printf, it prints to stdout and we want to capture this using the testing framework
// goal: test the actual printing by injecting/passing in the dependency of printing
// how: change the implementation to print to something we control

func TestGreet(t *testing.T) {
	// Buffer from the bytes package implements the Writer interface because it has the method Write(p []byte) (n int, err error)
	// Write requires a slice of bytes as input. Therefore we need to get bytes
	// https://yourbasic.org/golang/io-writer-interface-explained/
	buffer := bytes.Buffer{}
	Greet(&buffer, "you")

	got := buffer.String()
	want := "Hello, you"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// This opens up a web server session that is accessible over http://localhost:5000/
func TestAnotherWriterExample(t *testing.T) {
	AnotherWriterExample()
}