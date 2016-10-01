package main

import "fmt"

func main() {
	fact := factorial(50)
	fmt.Println("factorial:", fact)
}

func factorial(n int) uint64 {
	c := make(chan int)

	go func() {
		for i := 1; i <= n; i++ {
			c <- i
		}
		close(c)
	}()

	var val uint64 = 1
	for i := range c {
		//fmt.Println("range", i)
		val *= uint64(i)
	}

	return val
}

// 4! = 4*3*2*1
// Use channels to calculate factorials in order to make the program performant
// i.e. one thread is performing the loop of adding numbers to the channels
// and in parallel, the other thread is performing the multiplication calculation
