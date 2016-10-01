package main

import (
	"fmt"
	"sort"
)

type myints []int

func (p myints) Len() int {
	return len(p)
}

func (p myints) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p myints) Swap(i, j int) {
	holder := p[i]
	p[i] = p[j]
	p[j] = holder
}

func main() {
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	fmt.Println(n)
	sort.Sort(myints(n))
	fmt.Println(n)
}
