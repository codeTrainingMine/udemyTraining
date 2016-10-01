package main

import "fmt"

func main() {
	var x [58]string
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])
	for i := 65; i <= 122; i++ {
		//x[i-65] = string(i)
		x[i-65] = "hello"
	}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])

	//var b int
	//b = 10
	b := 10
	fmt.Println(b)

}
