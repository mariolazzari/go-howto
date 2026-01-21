package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	fmt.Printf("OS = %s\n", runtime.GOOS)
	fmt.Printf("ARCH = %s\n", runtime.GOARCH)

}
