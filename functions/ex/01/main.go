package main

import "fmt"

/*
Write a function which takes an integer.
The function will have two returns.
The first return should be the argument divided by 2.
The second return should be a bool that let’s us know whether or not the argument was even. For example:
half(1) returns (0, false)
half(2) returns (1, true)
*/

func twoReturnValue(x int) (int, bool) {
	var z1 int
	var z2 bool
	if x%2 == 0 {
		z1 = x / 2
		z2 = true
	} else {
		z1 = x / 2
		z2 = false
	}
	return z1, z2
	// or
	// return x/2, x%2 == 0
}

func main() {
	x, y := twoReturnValue(10)
	fmt.Println(x, y)
}
