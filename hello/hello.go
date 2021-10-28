package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "
const french = "French"
const frenchHelloPrefix = "Bonjour, "
const pigLatin = "Pig Latin"
const pigLatinHelloPrefix = "Ellohay, "

// Function name: Hello
// Returns: string
// Argument(s): string type called name, string type called language
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// Function name: greetingPrefix
// Returns: string type called prefix
// Specifying prefix string creates a variable prefix in the function with a zero value of "" for strings.
// Argument(s): string type called language
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case pigLatin:
		prefix = pigLatinHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	// You don't have to explicitly specify return prefix because it's defined in the function signature
	return
}

func main() {
	// Separate your "domain" code from the outside world (side-effects).
	// The fmt.Println is a side effect (printing to stdout) and the string we send in is our domain.
	fmt.Println(Hello("world", ""))
}