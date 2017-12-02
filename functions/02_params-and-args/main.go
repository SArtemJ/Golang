package main

import (
	"fmt"
)

func main() {
	greet("Bill")
	greet("Luke")
}

func greet(name string) {
	fmt.Println(name)
}
