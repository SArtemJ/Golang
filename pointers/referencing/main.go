package main

import (
	"fmt"
)

func main() {

	a := 55

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a


	fmt.Println(b)
	fmt.Println(*b)

	*b = 100+1
	fmt.Println(a)
	fmt.Println(*b)

}
