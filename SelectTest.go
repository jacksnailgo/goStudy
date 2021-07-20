package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		//  通过这种方式来模拟并行的协程执行（例如，RPC 操作）时造成的阻塞
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c3 <- "two"
	}()
	// 如果此时的 i是3  会提示一下报错  fatal error: all goroutines are asleep - deadlock!
	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("recieved", msg1)
		case msg2 := <-c2:
			fmt.Println("recieved", msg2)
		}
	}
}
