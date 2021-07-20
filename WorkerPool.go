package main

import (
	"fmt"
	"sync"
	"time"
)

/**
工作池 类比于Java的线程池，通过协程 和管道，实现了任务的并行处理
*/
func main1() {
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

	// 等待多个协程完成，我们可以使用 *wait group*

}

func gWorker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func waitWorker(id int, wg *sync.WaitGroup) {
	defer wg.Wait()
	fmt.Printf("worker %d starting \n\n", id)
	time.Sleep(time.Second)
	fmt.Printf("worker %d done \n", id)

}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go waitWorker(i, &wg)
	}
	wg.Wait()
}
