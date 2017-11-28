package main

import "fmt"

func main() {
	var smallNumber int
	var largeNumber int

	fmt.Println("Input small number")
	fmt.Scan(&smallNumber)
	fmt.Println("Input large number")
	fmt.Scan(&largeNumber)
	result := largeNumber%smallNumber
	fmt.Println("Result reminder is", result)
}
