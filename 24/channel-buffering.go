package main

import "fmt"

func main() {
	messages := make(chan string, 10)

	messages <- "bufferded"
	messages <- "channel"
	messages <- "1channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
