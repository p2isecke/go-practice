package iteration

import "strings"

// Function name: Repeat
// Returns: string
// Argument(s): string type called character, integer type called times
func Repeat(character string, times int) string {
	//return manual(character, times)
	// use the strings package instead of writing our own
	return strings.Repeat(character, times)
}

func manual(character string, times int)  string {
	// declare but don't initialize variable
	var repeated string

	for i := 0; i < times; i++ {
		// += adds the right operand to the left operand and assigns the result to the left operand
		repeated += character
	}
	return repeated
}

