package main

import (
	"fmt"
	"time"
)

func main() {
	var age int

	fmt.Println("How old are you?")
	fmt.Scan(&age)

	born_year := time.Now().Year() - age
	fmt.Printf("You are %d years old", born_year)
}
