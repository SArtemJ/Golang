package main

import "fmt"

func visit(numbers []int, callback func(int)) {
	for _, value := range numbers {
		callback(value)
	}
}

func main() {
	visit([]int{1,2,3,4,5}, func(value int) {
		fmt.Println(value)
	})
}