package main

import (
	"fmt"
	"time"
)

/**
工作池 类比于Java的线程池，通过协程 和管道，实现了任务的并行处理
*/
func main() {
	const nums = 5
	jobs := make(chan int, nums)
	results := make(chan int, nums)

	for w := 1; w <= 3; w++ {
		go gWorker(w, jobs, results)
	}

	for j := 1; j <= nums; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= nums; a++ {
		<-results
	}
}

func gWorker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}
