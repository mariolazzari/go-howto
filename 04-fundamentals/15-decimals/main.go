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
