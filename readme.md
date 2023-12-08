# Go Foundation

## Magesh Kuppan

## Schedule
- Commence      : 9:00 AM
- Tea Break     : 10:40 AM (20 mins)
- Lunch Break   : 12:30 PM (1 hour)
- Tea Break     : 3:00 PM (20 mins)
- Wind up       : 5:00 PM

## Methodology
- No powerpoints
- Discussion & Code

## Repository
- https://github.com/tkmagesh/cisco-go-dec-2023

## Software Requirements
- Go tools (https://go.dev/dl)
- Visual Studio Code (https://code.visualstudio.com)
- Go Extension for Visual Studio Code (https://marketplace.visualstudio.com/items?itemName=golang.Go)
### Verification
- >go version

## Why Go?
- Simplicity
    - ONLY 25 keywords
    - No access modifiers
    - No classes (only structs)
    - No inheritance (only composition)
    - No reference types (everything is a value)
    - No pointer arithmatic
    - No exceptions (only errors)
    - No try..catch..finally
    - No implicit type conversion
- Better Performance
    - Equivalent to C++
    - Close to the hardware (No VM)
        - Compile to machine code targetting the OS
        - Cross compilation is supported 
- Managed Concurrency
    - Light weight (1 goroutine = 4 KB whereas 1 OS Thread = 2 MB)
    - Language support for concurrency
        - go keyword, channel (data type), channel operator ( <- ), range & select case constructs
    - APIs 
        - sync package
        - sync/atomic package

- Reference
    - https://go.dev/talks/2015/go4cpp.slide

## Hello World

```
//01-hello-world.go

/* package declaration */
package main

/* import dependency packages */

/* package scope type/variable declarations */

/* package init function */

/* main function */
func main() {
	print("Hello World!")
}

/* other functions */
```
To run the program
> go run 01-hello-world.go

To create a build
> go build 01-hello-world.go

> go build -o [binary-name] 01-hello-world.go

Escape Analysis
> go build -gcflags="-m" [file_name.go]

### Cross Compilation
To get the environment variables
> go env

> go env GOOS GOARCH

To change the environment variables
> go env -w GOOS=[value] GOARCH=[value]

To get the list of supported platforms
> go tool dist list

To cross compile
> GOOS=windows GOARCH=amd64 go build 01-hello-world.go

## Data Types
- string
- bool

- int8
- int16
- int32
- int64
- int

- uint8
- uint16
- uint32
- uint64
- uint

- float32
- float64

- complex64 (real[float32] + imaginary[float32])
- complex128 (real[float64] + imaginary[float64])

- byte (alias for uint8)
- rune (alias for int32) (unicode code point)

## Variable Declarations
- using "var" keyword
- using ":="

## Scope
- function scope
    - can use ":=" for declaring and initializing variables
    - cannot have unused variables
- package scope
    - cannot use ":=" (use "var" instead)
    - can have unused variables

## Constants
- use "const" keyword
- constants can remain unused (both in function and package scope)

## Constructs
- if else
- for
- switch case

## Functions
- Can have return more than one result
- Variadic functions 
- Anonymous functions
- Higher Order Functions
    - Functions can be assigned as values to variables
    - Functions can be passed as arguments to other functions
    - Functions can be returned as return values from other functions

## Errors
- Errors are just values 
- Errors are not "thrown" but returned
- Typically, errors are objects implementing "error" interface
    - Error() method
- Create errors
    - errors.New()
    - fmt.Errorf()

## Collection
- Array
    - fixed sized typed collection
    - Members can be accessed using indexer
- Slice
    - varying sized typed collection
    - pointer to an underlying array
    - append() - to add new items
    - len() - # of values that can be accessed from the underlying array through the slice
    - cap() - memory initialized + memory uninitialized(only allocated)
![image slice](./images/slice.png)

- Map
    - typed collection of key/value pairs

## Modules & Packages
- Module
    - any code that need to be versioned and deployed together
    - typically a folder with go.mod file (manifest file)
        - module name
        - go runtime version
        - dependencies
    - To create a mod file
        > go mod init [module_name]
    - To execute a module
        > go run .
    - To create a build
        > go build .
        > go build -o [binary_name] .
- Package
    - internal organization of a module
    - typically folders
- Using 3rd party modules
    - To get 3rd party modules (downloaded GOPATH/pkg/mod folder)
        - > go get [module_name]
        - > ex: go get github.com/fatih/color
    - To update the go.mod file
        > go mod tidy
    - To download the documented (go.mod) dependencies
        > go mod download
    - To analyze the dependency graph
        > go mod graph
    - To get info about one dependency
        > go mod why [module_name]
    - To localize the dependencies
        > go mod vendor
    - To bypass the local vendor folder
        - > go run -mod=mod .
        - > go build -mod=mod .
    - To install a module as an executable (GOPATH/bin folder)
        > go install [module_name]
    - To verify the dependency modules
        > go mod verify
    - Reference: https://go.dev/ref/mod

## Structs
- Synonymous to classes
- Value types

## Methods
- functions with receivers
- Type definition & method definition must be in the same package