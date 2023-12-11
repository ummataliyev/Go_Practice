package main

import "fmt"

func main() {

	var history []string
	var modern_hisotry []string

	history = append(history, "Adolf Hitler")
	history = append(history, "Napoleon")
	history = append(history, "Stalin")

	modern_hisotry = append(modern_hisotry, "Elon Musk")
	modern_hisotry = append(modern_hisotry, "Bill Gates")
	modern_hisotry = append(modern_hisotry, "Zuck")

	historypop := pop(&history, 1)

	fmt.Println(history)
	fmt.Println(history[2])
	fmt.Println(historypop)

	fmt.Println("I want to talk with", history[1])
}
