package main

import (
	"fmt"
	"log"

	"example.com/err_greeting"

	"example.com/greetings"

	"example.com/slices_learning"
)

func main() {
	// Get a greeting message and print it.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err := greetings.Hello("Gladys")
	fmt.Println(message, err)
	//slice greeting
	message2, err2 := slices_learning.SliceGreet("Tim ")
	fmt.Println(message2, err2)

	greet_many()

	// Errors
	err_greeting.ErrHandling()
}

func greet_many() {
	names := []string{"Alice", "Dhruv", "Shan"}
	messages, errs := slices_learning.SliceGreetMultiPeople(names)
	if errs != nil {
		log.Fatal(errs)
	}
	fmt.Println(messages)
}
