package main

import "fmt"

func main() {

	x := 0
	incr := func() int {
		//func without name - anonymous
		x++
		return x
	}
	fmt.Println(incr())
	fmt.Println(incr())
}
