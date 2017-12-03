package main

import "fmt"

var x int

func increment() int {
	x++
	return x
}

func main() {
	//x++
	fmt.Println(increment())
	fmt.Println(increment())
}
