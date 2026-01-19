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

### How computer works

[Go proverbs](https://go-proverbs.github.io/)

### ASCII & UTF-8

```go
package main

import "fmt"

func main() {

	fmt.Println("Hello Gophers! ❤️💕😊👍😁(❁´◡`❁)£¥©🙌👌🎶😎🐼🦄🦁🐶😺🤓")

	// see emojis
	// Windows logo key + .

	/*
		ascii
		American Standard Code for Information Interchange
		2^8 (1 byte) = 256 unique values

		unicode
		2^32 (4 bytes) = 4,294,967,296 unique values
		more than enough to account for every character in every language

		utf-8
		(up to 4 bytes)
		stores unicode as binary
		If a character needs 1 byte that’s all it will use.
		If a character needs 4 bytes it will use 4 bytes.
		variable length encoding = efficient memory use
		common characters like “C” take 8 bits,
		rare characters like “💕” take 32 bits.
		Other characters take 16 or 24 bits.

		https://blog.hubspot.com/website/what-is-utf-8
		https://deliciousbrains.com/how-unicode-works/
	*/
}
```

### String literals

[String literals](https://go.dev/ref/spec#String_literals)

```go
package main

import "fmt"

func main() {

	fmt.Println("Hello Gophers! ❤️💕😊👍😁(❁´◡`❁)£¥©🙌👌🎶😎🐼🦄🦁🐶😺🤓")

	// raw string literal
	fmt.Println(`
	UTF-8 saves space.
	In UTF-8, common characters like “C” take 8 bits,
	while rare characters like “💕” take 32 bits.
	Other characters take 16 or 24 bits.
	A blog post like this one takes about
	four times less space in UTF-8
	than it would in UTF-32.
	So it loads four times faster.
	`)
}
```

## Hands-on exercise

### Printf

```go
package main

import "fmt"

func main() {
	fmt.Println("Ciao Mario, 👋")

	fmt.Println(`this is a
	 - raw
	 - string
	`)
}
```

## Fundamentals of Go

### Varaibles and zero values

```go
package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)

	b, c, d, _, f := 0, 1, 2, 3, "happiness"
	fmt.Println(b, c, d, f)

	// this would not work
	/*
		b, c, d, e := 0, 1, 2, 3
		fmt.Println(b, c, d)
	*/

	// this does work
	var g int
	fmt.Println(g)
	g = 43
	fmt.Println(g)

	// declare a variable to hold a VALUE of a certain TYPE
	// initializing a variable
	var h int = 44
	fmt.Println(h)

}
```

### Using printf with decimals

```go
package main

import (
	"fmt"
)

func main() {
	adams := 42
	fmt.Printf("42 as binary, %b \n", adams)
	fmt.Printf("42 as hexadecimal, %x \n", adams)

	// print these values as both binary & hexadecimal
	a, b, c, d, e, f := 0, 1, 2, 3, 4, 5

	fmt.Printf("a: %b %x\n", a, a)
	fmt.Printf("b: %b %x\n", b, b)
	fmt.Printf("c: %b %x\n", c, c)
	fmt.Printf("d: %b %x\n", d, d)
	fmt.Printf("e: %b %x\n", e, e)
	fmt.Printf("f: %b %x\n", f, f)

}
```

### Built-in types

[Types](https://pkg.go.dev/builtin)

### Scopes

[Scope](https://go.dev/ref/spec#Declarations_and_scope)
