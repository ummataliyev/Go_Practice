package main

import (
	"fmt"
	"math"
)

func main() {
	// Declare and initialize the variables
	greet := "Hello World!"
	news := "Assalomu alekum!"

	radius := 5
	pi := 3.14159
	surface := pi * math.Pow(float64(radius), 2)

	// Print the variables
	fmt.Println(greet)
	fmt.Println(news)
	fmt.Println("Radiusi", radius, "ga teng bo'lgan doira yuzi=", surface)
}
