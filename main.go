package main

import "fmt"

func main() {
	fmt.Println("thing")
	fmt.Println("another thing")
	thing()
}

func thing() {
	fmt.Println("Actually printing this time 2")
}
