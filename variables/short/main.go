package main

import "fmt"

func main()  {

	a := 10
	b := "go"
	c := 3.14
	d := true

	//%v value in default format
	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
}
