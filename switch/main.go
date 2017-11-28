package main

import "fmt"

func main() {

	switch "TestSwitch" {
	case "A":
		fmt.Println("this is a")
	case "B":
		fmt.Println("this is b")
	case "TestSwitch":
		fmt.Println("this is c")
	default:
		fmt.Println("Who is this?")
	}

}
