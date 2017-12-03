package main

import "fmt"

//Write a function with one variadic parameter that finds the greatest number in a list of numbers.

func findGreater(numbers ...int) int {
	var maxValue int
	for _, v := range numbers {
		if v > maxValue {
			maxValue = v
		}
	}
	return maxValue
}

func main() {
	findGreater := findGreater(4, 8, 1, 4, 76, 8, 9, 9)
	fmt.Println(findGreater)
}
