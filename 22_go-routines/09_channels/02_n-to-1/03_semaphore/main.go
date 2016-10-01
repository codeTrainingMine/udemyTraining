package main

import "fmt"

func main() {

	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		<-done
		<-done
		fmt.Println("before close(c)")
		close(c)
		fmt.Println("after close(c)")
	}()

	fmt.Println("before range loop")
	for n := range c {
		fmt.Println(n)
	}
}
