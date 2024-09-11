package main

import (
	"a5shawn.com/greetings"
	"fmt"
	"log"
)

func main() {
	hellos()
}

func hellos() {
	names := []string{"Shawn", "Jecity", "Gladys"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	for _, message := range messages {
		fmt.Println(message)
	}
}

func hello() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Jecity")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
