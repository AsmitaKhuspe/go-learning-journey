package main

import "fmt"

func getMarks() (marks int) {
	marks = 95
	return
}

func main() {
	result := getMarks()
	fmt.Println("Student's marks:", result)
}