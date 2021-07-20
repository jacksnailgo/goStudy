package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string, 2)

	go func() { messages <- "pingping" }()
	go func() { messages <- "good game" }()

	msg := <-messages
	fmt.Println(msg)
	fmt.Println(<-messages)

	done := make(chan bool, 1)
	go worker(done)
	res := <-done
	fmt.Println(res)
}

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}
