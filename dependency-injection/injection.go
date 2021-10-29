package dependency_injection

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// For the args passed into Greet, use the generic interface instead of specifcally bytes.Buffer
func Greet(writer io.Writer, name string) {
	// Use the writer to send the greeting to the buffer in our test
	// fmt.Fprintf takes an io.Writer to send the string to; fmt.Printf defaults to stdout
	fmt.Fprintf(writer,"Hello, %s", name)
}

// if Greet(writer bytes.Buffer), this wouldn't work because os.Stdout and bytes.Buffer are not the same
// they just implement the same interface
func generic() {
	Greet(os.Stdout, "Elodie")
}

// When you write an HTTP handler, you are given an http.ResponseWriter and the http.Request that was used to make the request.
// When you implement your server you write your response using the writer.
func handlerExample(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func AnotherWriterExample() {
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(handlerExample)))
}