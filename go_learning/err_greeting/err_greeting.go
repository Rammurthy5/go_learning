package err_greeting

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func ErrHandling() {
	// Get a greeting message and print it.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message2, err2 := greetings.Hello("")
	fmt.Println(message2, err2)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(message2)
}
