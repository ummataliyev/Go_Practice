package main

import "fmt"

func main() {
	var names []string

	names = append(names, "Muhammadali")
	names = append(names, "Fayzbek")
	names = append(names, "Shukurillo")

	fmt.Println("My friends names:", names)
	fmt.Println("Hi my friend", names[0])
}
