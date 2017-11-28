package main

import "fmt"

func main() {

	var resultSum int = 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 {
			resultSum += i
		} else if i%5 == 0 {
			resultSum += i
		}
	}
	fmt.Println(resultSum)
}
