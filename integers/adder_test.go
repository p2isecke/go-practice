package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		// Use %d to print integer instead of %q which prints string
		t.Errorf("expected '%d', got '%d'", expected, sum)
	}
}

// By specifying //OutputL 6, it causes the function to be compiled and executed from test context
func ExampleAdd(){
	sum := Add(1,5)
	fmt.Println(sum)
	// Output: 6
}