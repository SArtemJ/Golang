package main

import (
	"fmt"
)

func filter(numbers []int, callback func(int) bool) []int {
	var xs []int
	for _, value := range numbers {
		if callback(value) {
			xs = append(xs, value)
		}
	}
	return xs
}

func main() {
	xs := filter([]int{1,2,3,4,5}, func(value int) bool {
		return value > 1
	})
	fmt.Println(xs)
}
