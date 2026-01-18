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
