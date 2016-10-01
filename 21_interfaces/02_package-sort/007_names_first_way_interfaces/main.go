package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int {
	return len(p)
}

func (p people) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p people) Swap(i, j int) {
	holder := p[i]
	p[i] = p[j]
	p[j] = holder
}

func main() {
	studyGroup := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(people(studyGroup))
	sort.Sort(people(studyGroup))
	fmt.Println(studyGroup)
}
