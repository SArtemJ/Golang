package main

import "fmt"

type foo int

func main() {
	var myAge foo
	myAge = 27
	fmt.Printf("%T, %v \n", myAge, myAge)

	//02
	var yourAge int
	yourAge = 29
	fmt.Printf("%T, %v \n", yourAge, yourAge)

	//doesn't work
	//fmt.Println(myAge + yourAge)

	//conversion work
	fmt.Println(int(myAge) + yourAge)
}
