package main

import "fmt"

func main() {
	var x [55]string
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[50])
	for i := 68; i <= 122; i++ {
		x[i-68] = string(i)
	}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[50])
}
