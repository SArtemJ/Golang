package main

import "fmt"

func zeroFunc(x *int) {
	//x = 0
	//fmt.Printf("%p\n", &x)
	//fmt.Println(&x)
	*x = 0
}

func main() {
	x := 101
	//fmt.Printf("%p\n", &x)
	//fmt.Println(&x)
	zeroFunc(&x)
	fmt.Println(x)
}
