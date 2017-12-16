package main

import "fmt"

func main() {
	testOne()
	testTwo()
}
func testOne() {
	for i := 0; i < 35; i++ {
		fmt.Println("Test one - ", i)
	}
}

func testTwo() {
	for i := 0; i < 35; i++ {
		fmt.Println("Test two - ", i)
	}
}
