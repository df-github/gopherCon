package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func process(messages chan string, quit chan struct{}) {
	// Loop through your messages
	for m := range messages {
		// print the message with a for loop using range
		fmt.Println(m)
	}
	close(quit)
}

func main() {
	// declare the messages channel of type string and capacity of 5
	messagesChannel := make(chan string, 5)

	// declare a signal channel
	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	// launch process in a goroutine
	quit := make(chan struct{})
	go process(messagesChannel, quit)

	// declare 5 fruits in a []string
	fruits := []string{"Banana", "Ananas", "Apple", "Mango", "Strawberry"}

	// loop through the fruits and send them to the messages channel
	for _, f := range fruits {
		messagesChannel <- f
	}
	// close the messages channel
	close(messagesChannel)

	// wait for everything to finish
	<- quit

	fmt.Println("done")
}
