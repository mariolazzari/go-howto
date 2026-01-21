# Learn How To Code: Google's Go (golang) Programming Language

[Github](https://github.com/GoesToEleven/learn-to-code-go-version-03)

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

## Exercises

### A tour of Go

[A tour of Go](https://go.dev/tour/list)

### iota

[iota](https://go.dev/wiki/Iota)

### Measuring bits

```go
package main

import "fmt"

type ByteSize int

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	// ZB
	// YB
)

func main() {
	fmt.Printf("%d \t\t\t %b\n", KB, KB)
	fmt.Printf("%d \t\t %b\n", MB, MB)
	fmt.Printf("%d \t\t %b\n", GB, GB)
	fmt.Printf("%d \t\t %b\n", TB, TB)
	fmt.Printf("%d \t %b\n", PB, PB)
	fmt.Printf("%d \t %b\n", EB, EB)
}

/*
PB		    1125899906842624
EB		 1152921504606846976
int		18446744073709551615
*/
```

### Zero values

```go
package main

import "fmt"

var y int

func main() {
	fmt.Println(y)

	z := 42
	fmt.Println(z)

	a, b := 43, "Happiness"
	fmt.Println(a, b)

	var c float32 = 42.42
	fmt.Printf("%v is of this type %T\n",c, c)

	e, f, _ := 45, 46, 47
	fmt.Println(e, f)

}

/*
var for zero value
short declaration operator
multiple initializations
var when specificity is required
blank identifier

*/
```

### Show values and types

```go
package main

import "fmt"

func main() {
	s, i, f := "Happiness", 42, 42.42
	fmt.Printf("%v is of type %T\n",i,i)
	fmt.Printf("%v is of type %T\n",f,f)
	fmt.Printf("%v is of type %T\n",s,s)
}
```

### Printf

```go
package main

import "fmt"

func main() {
	x, y, z := 747, 911, 90210

	fmt.Printf("%d \t\t %b \t\t %#X\n",x,x,x)
	fmt.Printf("%d \t\t %b \t\t %#X\n",y,y,y)
	fmt.Printf("%d \t\t %b \t %#X\n",z,z,z)
}
```

### Signed and unsigned

```go
package main

import "fmt"

func main() {
	var x uint8 = 255
	var y int8 = 127

	fmt.Printf("%v is of type %T\n", x, x)
	fmt.Printf("%v is of type %T\n", y, y)
}
```

## Programming fundamentals

### Scope

[scope](https://go.dev/ref/spec#Declarations_and_scope)

```go
package main

// the imported package "fmt"
// is in the "file block" scope
// of this file
import "fmt"

// x is in "package block" scope
var x = 42

func main() {
	// x can be accessed here
	fmt.Println(x)

	// x can be accessed here
	printMe()

	// y does not exist here yet
	// so this won't work
	// fmt.Println(y)

	// y is in "block scope"
	y := 43
	fmt.Println(y)

	p1 := person{
		"James",
	}
	p1.sayHello()

	// variable "shadowing"
	x := 32
	fmt.Println(x)

	// package block x is still the same
	printMe()

	//also good to know

	if z := 82; z > 50 {
		fmt.Println(z)
	}
	// you can't access z here
	// fmt.Println(z)
	/*
		take a look at the "comma ok idiom"
		https://go.dev/doc/effective_go#maps
	*/
}

func printMe() {
	//x can be accessed here
	fmt.Println(x)
}

// type person is in "package block" scope
type person struct {
	first string
}

// the method sayHello
// which is attached to VALUES of TYPE person
// is in "package block" scope
func (p person) sayHello() {
	fmt.Println("Hi, my name is", p.first)
}
```

### Terminal commands

[Linux](https://ss64.com/bash/)

### Github ssh

```sh
ssh-keygen -t ed25519 -C "your_email@example.com"
```

## Dev enviroment
