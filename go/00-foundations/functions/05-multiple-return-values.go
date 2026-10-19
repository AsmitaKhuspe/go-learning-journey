package main

import "fmt"

func getNumbers() (int, int) {
	return 10, 20
}

func main() {
	first, second := getNumbers()

	fmt.Println("First number:", first)
	fmt.Println("Second number:", second)
}