package main

import "fmt"

func main() {
	fmt.Print("Please enter the following information:\n")

	fmt.Print("Street: ")
	street := getUserInput()

	fmt.Print("Region: ")
	region := getUserInput()

	fmt.Print("City: ")
	city := getUserInput()

	fmt.Print("Country: ")
	country := getUserInput()

	result := fmt.Sprintf("%s street, %s region, %s city, %s country", street, region, city, country)
	fmt.Println(result)
}

func getUserInput() string {
	var input string
	fmt.Scanln(&input)
	return input
}
