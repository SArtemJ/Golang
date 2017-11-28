package main

import "fmt"

func main() {
	var someName string = "Name"
	fmt.Println("Hello some ", someName)
	fmt.Scan(&someName)
	fmt.Println("Hello some ", someName)

}
