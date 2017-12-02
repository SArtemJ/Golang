package main

import "fmt"

func main() {
	greet("John", "Doe")
}

func greet(firstName string, secondName string) {
	fmt.Println(firstName, secondName)
}
