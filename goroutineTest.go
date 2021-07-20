package main

import (
	"fmt"
	"time"
)

func main() {
	ff("direct")
	go ff("goroutine")
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}

func ff(ss string) {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
