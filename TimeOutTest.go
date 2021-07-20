package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case result := <-c1:
		fmt.Println(result)
	// 超时处理
	case <-time.After(1 * time.Second):
		fmt.Println("timeOut 1")
	}
	// 非阻塞通道操作
	messages := make(chan string)
	signals := make(chan bool)
	// 非阻塞接受
	select {
	case msg := <-messages:
		fmt.Println("recieved message", msg)
	default:
		fmt.Println("no message recieved")
	}
	// 非阻塞发送
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
		// 多路非阻塞选择器
		select {
		case msg := <-messages:
			fmt.Println("recieved message", msg)
		case sig := <-signals:
			fmt.Println("recieved signal", sig)
		default:
			fmt.Println("no activity")

		}

	}

}
