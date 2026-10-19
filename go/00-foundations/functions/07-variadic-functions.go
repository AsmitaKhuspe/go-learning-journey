package main

import "fmt"

func greet(names ...string) {
	fmt.Println(names)
}

func main() {
	greet("Asmita")
	greet("Asmita", "Rahul")
	greet("Asmita", "Rahul", "Priya")
}