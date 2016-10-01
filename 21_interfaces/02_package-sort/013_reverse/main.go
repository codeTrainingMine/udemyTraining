package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Zeno", "John", "Al", "Jenny"}

	r := sort.Reverse(sort.StringSlice(s))
	// cannot do this because *sort.reverse is not exported:
	//var *sort.reverse r = sort.Reverse(sort.StringSlice(s))
	fmt.Printf("%T \n", r)
	fmt.Printf("%T \n", s)

	sort.Sort(r)
	sort.Sort(sort.StringSlice(s))
	fmt.Println(r)
	fmt.Println(s)
}
