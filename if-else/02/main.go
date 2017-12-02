package main

import "fmt"

func main()  {

	b := true

	//init variable in if statement before check
	if some := "Bread"; b {
		fmt.Println(some)
	}
	//fmt.Println(some) - can't use some here
}
