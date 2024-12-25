package main

import (
	"fmt"
)

func main() {
	str := "Hello, World!"
	p := &str
	*p = "Hello, Go!"

	fmt.Println("Value of str is: and pointer address is: ", str, *p)
}