package main

import (
	"fmt"
)

func main() {

	a := 55

	fmt.Println("Value a is ", a)
	fmt.Println("Address a is - ", &a)

	var b *int = &a


	fmt.Println("Pointer to a is - ", b)
	fmt.Println("value from pointer is - ", *b)

	*b = 100+1
	fmt.Println(a)
	fmt.Println(*b)

}
