package main

import "fmt"

func main() {
    marks := 82

    if marks >= 90 {
        fmt.Println("Excellent")
    } else if marks >= 75 {
        fmt.Println("Good")
    } else {
        fmt.Println("Keep practicing")
    }
}