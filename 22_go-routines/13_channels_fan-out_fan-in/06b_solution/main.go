package main

import (
	"fmt"
	"math/rand"
)

func main() {

	in := gen()

	f := factorial(in)

	results := make(chan int)
	// merge
	go func() {
		for _, n := range f {
			go func(c chan int) {
				for i := range c {
					results <- i
				}
			}(n)
		}
	}()

	// print
	for {
		fmt.Println(<-results)
	}
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			out <- rand.Intn(13) + 1
		}
		close(out)
	}()
	return out
}

func factorial(in <-chan int) []chan int {
	out := make([]chan int, 3)
	go func() {
		i := 0
		for n := range in {
			out[i] = make(chan int)
			go func(n int, c chan int) {
				c <- fact(n)
				close(c)
			}(n, out[i])
			i++
		}
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
