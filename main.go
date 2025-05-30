package main

import (
	"fmt"

	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"gorm.io/gorm"
)

func main() {
	_, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("could not connect to database")
	}
	fmt.Println("thing")
	fmt.Println("another thing")
	thing()
}

func thing() {
	fmt.Println("Actually printing this time 2")
}
