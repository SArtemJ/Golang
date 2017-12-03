package main

import "fmt"

func main() {
	mySlice := make([]int, 0, 5)
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))

	for i := 0; i < 80; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len: ", len(mySlice), "Cap: ", cap(mySlice), "value: ", mySlice[i])
	}
}
