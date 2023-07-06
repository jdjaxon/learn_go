package main

import (
	"fmt"
	"sync"
)

func worker(tpool chan int, wg *sync.WaitGroup) {
	for p := range tpool {
		fmt.Println(p)
		wg.Done()
	}
}

func main() {
	tpool := make(chan int, 100)
	var wg sync.WaitGroup
	for i := 0; i < cap(tpool); i++ {
		go worker(tpool, &wg)
	}

	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		tpool <- i
	}

	wg.Wait()
	close(tpool)
}
