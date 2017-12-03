package main

import (
	"fmt"
)

//func greeting() {
//	fmt.Println("Hello world")
//}

func main() {

	greeting := func() {
		fmt.Println("Hello world")
	}

	greeting()
	fmt.Printf("%T\n", greeting)
}
