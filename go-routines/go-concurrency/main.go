package main

import "fmt"

func main() {

	//go routine 1 - main
	//go routine 2 - testOne
	go testOne()
	//go routine 3 - testTwo
	go testTwo()
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
