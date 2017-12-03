package main

import "fmt"

func main() {
	fmt.Println((true && false) || (false && true) || !(false && false))
	//(true && false)  - false
	//(false && true) - false -> false || false = false
	//(false && false) - false -> !false = true
	// false || true = true
}
