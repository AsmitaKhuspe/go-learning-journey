package main

import "fmt"

func introduce(name string, age int) {
	fmt.Println("My name is", name)
	fmt.Println("My age is", age)
}

func main() {
	introduce("Asmita", 23)
}