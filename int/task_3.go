package main

import "fmt"

func main() {
	var a, b float64

	fmt.Print("Enter the first number: ")
	fmt.Scan(&a)

	fmt.Print("Enter the second number: ")
	fmt.Scan(&b)

	// plus
	result1 := a + b
	fmt.Printf("Result of a + b = %f\n", result1)

	// minus
	result2 := a - b
	fmt.Printf("Result of a - b = %f\n", result2)

	// multiplication
	result3 := a * b
	fmt.Printf("Result of a * b = %f\n", result3)

	// division
	if b != 0 {
		result4 := a / b
		fmt.Printf("Result of a / b = %f\n", result4)
	} else {
		fmt.Println("Cannot divide by zero.")
	}

	// remainder (modulo)
	if b != 0 {
		result5 := int64(a) % int64(b) // Modulo operation requires integers
		fmt.Printf("Result of a %% b = %d\n", result5)
	} else {
		fmt.Println("Cannot calculate remainder when the second number is zero.")
	}
}
