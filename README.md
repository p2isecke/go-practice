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

