package main

import (
	"practice1/greetings"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	var message string
	var err error
	message, err = greetings.Greetings("Gladys")
	
	if err != nil {
		log.Fatal(err)
	}

	log.Println(message)
}

