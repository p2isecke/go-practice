# Notes from [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/)

## Terminology
* **Go Modules:** handles dependencies, versions, and allows you to run code outside of `GOPATH`

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
More info: https://golang.org/doc/modules/gomod-ref
To use **modules**:
1. Select a directory outside of `GOPATH`. This is your new project root.
3. `go.mod` file generated containing:
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

### Functions
* Function names:
  * Public functions start with **capital letter**
  * Private functions start with **lowercase letter**
* Function signatures
  * When you have more than one argument of the same type you can shorten
    * Example: (x int, y int) --> (x, y int)
  * Named return values are included in documentation. More info: https://github.com/golang/go/wiki/CodeReviewComments#named-result-parameters

### Initializing Variables
* More info: https://gobyexample.com/variables 
* Declare and initialize variable: `:=`
* Declare only: `var <name> <type`
* Add and assign: `+=`

### Testing
* Filename: `xxx_test.go`
* Test function prefixes with `Test`
* Test function takes in one argument: `t *testing.T`
  * To use `t *testing.T`, `import "testing"`
  * `t` of type `*testing.T` hooks into the testing framework
* Use subtests to group tests around the same thing

### Benchmarking
* Measures how long code takes to execute by running it b.N times
* By default, benchmarks are run **sequentially**

