package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen()

	c := fanOut(in)

	for n := range merge(c) {
		fmt.Println(n)
	}
}

func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for i := 1; i < 6; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func fanOut(in chan int) []chan int {
	out := make([]chan int, 5)
	var wg sync.WaitGroup

	var i int
	for n := range in {
		wg.Add(1)
		go func(n int, ch chan int) {
			ch <- fact(n)
			close(ch)
			wg.Done()
		}(n, out[i])
		i++
	}

	go func() {
		wg.Wait()
	}()

	return out
}

func merge(cs []chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		go func(ch chan int) {
			for n := range ch {
				out <- n
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
