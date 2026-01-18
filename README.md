# Learn How To Code: Google's Go (golang) Programming Language

## Course introduction

### Course resources

- [Forum](https://forum.golangbridge.org/)
- [Documentation](https://go.dev/doc/)
- [Specification](https://go.dev/ref/spec)
- [Effective Go](https://go.dev/doc/effective_go)
- [Standard library](https://pkg.go.dev/std)
- [Blog](https://go.dev/blog/)
- [Modules](https://go.dev/blog/using-go-modules)
- [Modules reference](https://go.dev/ref/mod)
- [How to Write Go Code](https://go.dev/doc/code)
- [Create a Go module](https://go.dev/doc/tutorial/create-module)
- [How to Use Go Modules](https://www.digitalocean.com/community/tutorials/how-to-use-go-modules)
- [Go modules](https://www.practical-go-lessons.com/chap-17-go-modules)
- [Go proverbs](https://go-proverbs.github.io/)
- [Go by Example](https://gobyexample.com/)
- [Naming a module](https://go.dev/doc/modules/managing-dependencies#naming_module)
- [Gophers](https://github.com/ashleymcnamara/gophers)
- [Gophers](https://github.com/egonelbre/gophers)
- [Leran to code Go](https://github.com/GoesToEleven/learn-to-code-go-version-03)
- [A tour of Go](https://go.dev/tour/welcome/1)

## Getting going

### Why Go?

Go is a statically typed, compiled language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. It is known for its simplicity, efficiency, and concurrency features. Go's syntax is clean and easy to read, making it a great choice for both beginners and experienced developers.

### First program

```go
// main package
package main

import (
	"fmt"
)

// main entry point
func main() {
	fmt.Println("Hello, Go!")
}
```

### Format printing

[fmt](https://pkg.go.dev/fmt)
[Printf](https://pkg.go.dev/fmt#example-Printf)
[formats](https://pkg.go.dev/fmt#hdr-Printing)

```go
package main

import (
	"fmt"
)

func main() {
	const name, age = "Kim", 22
	n, err := fmt.Printf("%s is %d years old.\n", name, age)

	// It is conventional not to worry about any
	// error returned by Printf.
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	fmt.Printf("Bytes written: %d\n", n)

	fmt.Printf("name type is %T and its value is %v", name, name)
}
```
