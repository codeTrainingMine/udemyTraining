package main

import (
	"fmt"
	"sync"
)

var count int64
var wg sync.WaitGroup

func main() {
	count := make(chan int)
	go incrementor("1", count)
	go incrementor("2", count)

	//for n := range count {
	//for {
	var final int
	for i := 1; i <= 40; i++ {
		//fmt.Println("main", <-count)
		<-count
		final = i
	}

	fmt.Println("Final Counter:", final)
}

func incrementor(s string, out chan int) {
	for i := 0; i < 20; i++ {
		out <- i
		fmt.Println("Process: "+s+" printing:", i)
	}
}
