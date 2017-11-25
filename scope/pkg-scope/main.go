package main

import "fmt"

var someX = 27

func main() {

	fmt.Println(someX)
	secondFunc()

	t := 101
	fmt.Println(t)
}

func secondFunc() {
	fmt.Println(someX)
	//fmt.Println(t)
}
