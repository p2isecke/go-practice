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
More information: https://golang.org/doc/modules/gomod-ref
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
* `main` package  with `main` function

### Testing
* Filename: `xxx_test.go`
* Test function prefixes with `Test`
* Test function takes in one argument: `t *testing.T`
  * To use `t *testing.T`, `import "testing"`
  * `t` of type `*testing.T` hooks into the testing framework
* Use subtests to group tests around the same thing

