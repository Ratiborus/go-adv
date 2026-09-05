package main

import (
	"fmt"
	"math/rand"
)

func main() {
	intCh := make(chan int)
	go func(c chan int) {
		slice := make([]int, 0, 10)
		for range 10 {
			slice = append(slice, generate())
		}
		for _, v := range slice {
			c <- v
		}
		close(c)
	}(intCh)
	sqCh := make(chan int)
	go func(source chan int, target chan int) {
		for num := range source {
			sqCh <- num * num
		}
		close(sqCh)
	}(intCh, sqCh)

	for sq := range sqCh {
		fmt.Printf("%v ", sq)
	}
}

func generate() int {
	return rand.Intn(100)
}
