# Notes from [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/)

## Go Environment
* `go 1.11` introduces **modules**
* `go 1.16` made **modules** the default
  * `GOPATH` is no longer recommended
* Run `go fmt` on every file save.
* You can run Go documentation locally:
  * `go get golang.org/x/tools/cmd/godoc`
  * `godoc -http :8000`
  * http://localhost:8000/pkg 

### Go Modules
* Handles dependencies, versions, and allows you to run code outside of `GOPATH`
* More info: https://golang.org/doc/modules/gomod-ref
* To use **modules**:
  1. Select a directory outside of `GOPATH`. This is your new project root.
  2. `go.mod` file generated containing:
     1. Module path
     2. Go version
     3. Dependency requirements
     4. Indicate where code is available for download
     5. Example `go.mod` file:
        ```
        ❯ cat go.mod 
        module learn-go-with-tests
    
        go 1.16
        ```
   
## Code Structure/Best Practices
* Packages group related `go` code together
  * Go source files can only have one package per directory
* `main` package  with `main` function
* Consider adding examples in the `_test.go` file. This is preferred over including examples in READMEs since they are compiled and tested.
  * Run `go test -v` to see the output.
* Formatting: https://pkg.go.dev/fmt
* Starts with **lowercase** --> Private and not accessible outside the package it's defined in
* When you call a function or a method the arguments are copied.
* `var` lets you define values global to the package
* Constants are like variables but can't be modified after declaration

### Dependency Injection
* Allows you to control where your data is written by injecting a dependency
* Enables:
  * Testing, especially if dependencies are tightly coupled with some code, e.g database connection pool
  * Decoupling where data goes from how to generate it, e.g. function generates data and writes to a db
  * Reusability in other contexts

Example: [fmt.Printf:](https://pkg.go.dev/fmt#Printf) [source code](https://cs.opensource.google/go/go/+/refs/tags/go1.17.2:src/fmt/print.go;l=212)
```
// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func Printf(format string, a ...interface{}) (n int, err error) {
    // Printf calls Fprintf with os.Stdout as an arg
    // Fprintf expects an io.Writer as a first argument. Here we pass in os.Stdout
    // Writer is an interface
    // Therefore os.Stdout implements io.Writer
	return Fprintf(os.Stdout, format, a...)
}
```
[Fprintf:](https://cs.opensource.google/go/go/+/refs/tags/go1.17.2:src/fmt/print.go;drc=refs%2Ftags%2Fgo1.17.2;l=202)
```
// Fprintf formats according to a format specifier and writes to w.
// It returns the number of bytes written and any write error encountered.
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	p := newPrinter()
	p.doPrintf(format, a)
	n, err = w.Write(p.buf)
	p.free()
	return
}
```
[io.Writer:](https://cs.opensource.google/go/go/+/master:src/io/io.go;l=96?q=io.Writer&ss=go%2Fgo)
```
// Writer is the interface that wraps the basic Write method.
//
// Write writes len(p) bytes from p to the underlying data stream.
// It returns the number of bytes written from p (0 <= n <= len(p))
// and any error encountered that caused the write to stop early.
// Write must return a non-nil error if it returns n < len(p).
// Write must not modify the slice data, even temporarily.
//
// Implementations must not retain p.
type Writer interface {
	Write(p []byte) (n int, err error)
}
```

### Functions
* Function names:
  * Public functions start with **capital letter**
  * Private functions start with **lowercase letter**
* Function signatures
  * When you have more than one argument of the same type you can shorten
    * Example: (x int, y int) --> (x, y int)
  * Named return values are included in documentation. More info: https://github.com/golang/go/wiki/CodeReviewComments#named-result-parameters
* You can ignore returned values by using the blank identifier `_`
* **Variadic functions:** take in any number of arguments
  * `func functionName(things ...[]type)`
* `reflect.DeepEqual` which is useful for seeing if any two variables are the same
  * `reflect.DeepEqual` is not "type safe"
* You cannot overload a function like:
  ```
  func Area(circle Circle) float64 { ... }
  func Area(rectangle Rectangle) float64 { ... }
  ```
  * You can however:
    * Have functions with the same name in _differen_ packages
    * Define **methods** on the type/struct

### Methods
* Method is a function **with a receiver**
* Invoked on an instance of a particular type
  * In Java, you use `this.`
* `func (receiverName ReceiverType) MethodName(args)`
* Function vs Method:
  * Function: Area(Rectangle)
  * Method: rectangle.Area()
* **Best Practice:** Receiver variable is the first letter of the type

### Structs
* Custom types
* Let's you bundle related data together
* More info: https://golang.org/ref/spec#Interface_types
* Anonymous struct: declare a slice of structs using []struct with fields
* Struct pointers are automatically derefenced so you don't have to do `(*r)`
  * More info: https://golang.org/ref/spec#Method_values

### Interfaces
* Allows you to define functions that can be used with different types
* In Java, "my type foo implements interface bar"
* Interface resolution is _implicit_. If the type you pass in matches what the interface is asking for, it will compile.
* By declaring an interface, the helper is decoupled from the concrete types and only has the method it needs to do its job.
* Hides complexity from other parts of the system
* Can be nilable
  * If you try to access a value that is nil, it will through a `runtime panic`

### Types
* Go lets you create new types from existing ones:
  * `type MyName OriginalType`
* You can declare methods on these new types

### Error Handling
* Return an error from your function so that other code can do something with it
* `errors.New("message")` creates a new error with any message
* You can use a tool called errcheck (`go get -u github.com/kisielk/errcheck`) to make sure that all errors are handled
* Don't just check for errors, handle them: https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully
* Errors can be converted to a string with `.Error()`
* Constant errors: https://dave.cheney.net/2016/04/07/constant-errors

### Pointers
* Memory address (pointer): `&variableName`
* Pass by reference
* Go copies the values when you pass them to functions/methods
  * If you need to change state you'll need to pass a pointer to the thing you want to change
* When you pass a pointer to a function, you're passing the address
  * You can't change the original unless you pass the pointer
* When you don't pass a pointer, you pass a copy
* Can be `nil`
  * Be sure to check if it's nil to avoid runtime exceptions

### Initializing Variables
* More info: https://gobyexample.com/variables 
* Declare and initialize variable: `:=`
* Declare only: `var <name> <type`
* Add and assign: `+=`

### Arrays
* Arrays have **fixed capacity**, defined when you declare it
* `myArray := [5]int{1,2,3,4,5}` --> arrayName := [N]type{values,...}
* `myArray := [...]int{1,2,3,4,5}` --> arrayName := [...]type{values,...}

### Slices
* Do not have fixed capacity, encode the size of the collection
  * While it's not fixed there's still capacity. You can't index out of bounds, otherwise you will get an error:
  `panic: runtime error: index out of range [10] with length 2`
* `mySlice := []int{1,2,3,4,5}` --> sliceName := [N]type{values,...}
* `mySlice := make([]type, starting capacity)`
* You can get a portion of slices: slice[low:high] 
  * `[1:]`: from 1 to the end
* Make copies of slices before modifying it

### Maps
* Store values using a key
* `map[keyType]valueTYpe`
* Keys can only be a comparable type
  * More info: https://golang.org/ref/spec#Comparison_operators
* Values can be of any type, including another map
* 2 return values: 1.the lookup value and 2. a boolean indicating if it was found
* You can modify maps without passing an address
* When you pass a map to a function/method, you're copying it, but just the pointer, not the underlying data structure
* Can be `nil`; haves like an empty map when reading
* Writing to a `nil` map will cause runtime panic
* **Never** initialize an empty map variable: `var m map[string]string`
  * Instead create an empty hash map and point to it:
    * `var m = make(map[string]string` or
    * `var m = map[string]string{}`
* Does not throw an error if the value exists; it overwrites it

### Testing
* Filename: `xxx_test.go`
* Test function prefixes with `Test`
* Test function takes in one argument: `t *testing.T`
  * To use `t *testing.T`, `import "testing"`
  * `t` of type `*testing.T` hooks into the testing framework
* Use subtests to group tests around the same thing
* Table driven tests: https://github.com/golang/go/wiki/TableDrivenTests
  * Useful when testing various implementations of an interface, varying data being passed in
  * Test speaks to us more clearly, as if it were an assertion of truth, not a sequence of operations


### Benchmarking
* Measures how long code takes to execute by running it b.N times
* By default, benchmarks are run **sequentially**
