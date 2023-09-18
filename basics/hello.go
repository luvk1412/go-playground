package main

import (
	"fmt"
	"unsafe"
)

func swap(a, b string) (string, string){
	return b, a
}

func main() {
	fmt.Println(swap("Hello Luv", "Kumar"))
	fmt.Println(unsafe.Sizeof(complex128(0)))
}