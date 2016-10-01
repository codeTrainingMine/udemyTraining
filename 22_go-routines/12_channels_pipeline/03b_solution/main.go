package main

import "fmt"

func main() {
	// Set up the pipeline and consume the output.
	for n := range factorial(gen(100)) {
		fmt.Println(n)
	}
}

func gen(n int) chan int {
	out := make(chan int)
	go func() {
		for i := 1; i <= n; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func factorial(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			total := 1
			for i := n; i > 0; i-- {
				total *= i
			}
			out <- total
		}
		close(out)
	}()
	return out
}

// What realizations did you have about working with concurrent code when building your solution?
// 	1. tricky and counter-intuitive
// 	2. program a bit harder to read and understand
// 	3. must have a go func for "queuing the work" and another go func for "de-queuing and processing the work"
// What insights did you have which helped you understand working with concurrency?
// 	1. the "printing loop" should also be the one that sends the code off for processing, otherwise code could get complicated
//	2. and easy to make mistake, is to put the factorial worker in a loop which calls it via a go func. this would produce multiple out channels which is not what we want
//  3. it's easier to base this on a working bit of code (i.e. the sq example)
//  4. inversely (to #3) it's harder to come up with this pattern by yourself
//	5. this seems to run quite fast when done in parallel!
//	6. watch out for i/n for-loop index confusion i.e "for i := n; i > 0; n--" leads to infinite loop!
// https://goo.gl/uJa99G
