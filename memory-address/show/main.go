package main

import "fmt"

func main() {
	x := 101
	fmt.Println("x - ", x)
	fmt.Println("x address is - ", &x)
	fmt.Printf("%d \n", &x)
}
