package main

import "fmt"

func main() {

	myGreeting := map[int]string{
		0: "Hi",
		1: "Hello",
		2: "Aloha",
		3: "Bonjour",
	}

	for key, a := range myGreeting {
		fmt.Println(key, " - ", a)
	}
}
