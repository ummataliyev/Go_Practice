package main

import (
	"fmt"
	"strings"
)

func main() {
	street := "London"
	region := "Piwerma"
	counrty := "Great Britain"

	result := fmt.Sprintf("%s street, %s region, %s country", street, region, counrty)

	fmt.Println(strings.ToUpper(result))
	fmt.Println(strings.ToLower(result))
	fmt.Println(strings.Title(result))
}
