package main

import "fmt"

func main() {
	x := 55
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "Message in closure"
		var z = "test"
		fmt.Println(y)
		fmt.Println(z)
	}
	//fmt.Println(z)
	//fmt.Println(y) --not work because y declare in closure {}
}
