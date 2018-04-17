package main

import "fmt"

func main() {
	messages := make(chan string)

	for i := 0; i < 3; i++ {
		go func() { messages <- "ping" }()
	}

	msg := <-messages
	fmt.Println(msg)
}
