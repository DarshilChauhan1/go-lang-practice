package main

import (
	"fmt"
)

type person struct {
	name   string
	age    int16
	height float64
}

func main() {
	person1 := person{name: "John", age: 25, height: 5.8}
	person2 := person{name: "Jane", age: 22, height: 5.5}

	person1.name = "Doe"
	fmt.Printf("Person1: %v\n", person1)
	fmt.Printf("Person2: %v\n", person2)
}
