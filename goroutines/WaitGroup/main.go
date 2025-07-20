package main

import (
	"fmt"
	"sync"
	"time"
)

func pn(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(n)
	time.Sleep(1 * time.Second)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go pn(i, &wg)
	}

	wg.Wait()

	fmt.Println("all goroutine is finished")
}
