package main

import "fmt"

func main() {
	var student []string
	var students [][]string
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)

	student = append(student, "aaa")
	students = append(students, student)
	students = append(students, student)

	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)
}
