package main 

import (
	"practice2/Hello"
	"log"
	"fmt"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := Hello.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	for _, message := range messages {
		fmt.Println(message)
	}
}

