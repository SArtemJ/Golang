package main

import "fmt"

func main() {
	x := 100
	changeMe(x)
	fmt.Println(x)

}

func changeMe(y int)  {
	fmt.Println(y)
	y = 101
	fmt.Println(y)
}
