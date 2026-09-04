package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	five()
}

func one() {
	numbers := []int{10, 20, 30, 40, 50}
	buf := make(chan int)
	go func() {
		for _, val := range numbers {
			buf <- val
		}
	}()
	sum := 0
	for range numbers {
		val := <-buf
		sum += val
	}
	fmt.Println(sum)
}

func two() {
	tasks := []string{"task-A", "task-B", "task-C", "task-D"}
	var wg sync.WaitGroup
	buf := make(chan string, len(tasks))
	for _, str := range tasks {
		wg.Add(1)
		go func(s string) {
			buf <- fmt.Sprintf("done: %s", s)
			wg.Done()
		}(str)
	}
	wg.Wait()
	close(buf)
	for s := range buf {
		fmt.Println(s)
	}
}

func three() {
	resultMsg := "result: ok"
	errorMsg := "error: something went wrong"

	resCh := make(chan string)
	errCh := make(chan string)

	go func() {
		resCh <- fmt.Sprintf("res: %s", resultMsg)
	}()
	go func() {
		errCh <- fmt.Sprintf("err: %s", errorMsg)
	}()
	for range 2 {
		select {
		case res := <-resCh:
			fmt.Println(res)
		case err := <-errCh:
			fmt.Println(err)
		}
	}
}

func four() {
	message := "async task executed"
	go func() {
		fmt.Println(message)
	}()
	time.Sleep(time.Millisecond * 100)
}

func five() {
	numbers := []int{2, 3, 4, 5, 6}
	resCh := make(chan int, len(numbers))
	for _, val := range numbers {
		go func(v int) {
			resCh <- v * v
		}(val)
	}
	sum := 0
	for range len(numbers) {
		res := <-resCh
		sum += res
	}
	fmt.Println(sum)
}
