package main

import "fmt"

func main() {
	x := 45
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "Mr. Anderson"
		fmt.Println(y)
	}
	//fmt.Println(y) //outside scope of y

}
