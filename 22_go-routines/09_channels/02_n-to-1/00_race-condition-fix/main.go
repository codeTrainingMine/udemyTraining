package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)

	var wg sync.WaitGroup
	var mt sync.Mutex

	go func() {
		mt.Lock()
		wg.Add(1)
		mt.Unlock()
		for i := 0; i < 10; i++ {
			c <- i
		}
		mt.Lock()
		wg.Done()
		mt.Unlock()
	}()

	go func() {
		mt.Lock()
		wg.Add(1)
		mt.Unlock()
		for i := 0; i < 10; i++ {
			c <- i
		}
		mt.Lock()
		wg.Done()
		mt.Unlock()
	}()

	go func() {
		mt.Lock()
		wg.Wait()
		mt.Unlock()
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
