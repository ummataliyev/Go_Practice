package main

import (
	"fmt"
	"math"
)

func main() {
	var x int

	fmt.Println("Enter the number:")
	fmt.Scan(&x)

	fmt.Printf("%d squared is %d\n", x, int(math.Pow(float64(x), 2)))
	fmt.Printf("%d cubed is %d\n", x, int(math.Pow(float64(x), 3)))

}
