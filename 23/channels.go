package main

import "fmt"

func main() {
	messages := make(chan string)
	go func() {
		messages <- "ping"
		messages <- "ping1"
	}()

	msg := <-messages
	fmt.Println(msg)
}
