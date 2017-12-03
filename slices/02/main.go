package main

import "fmt"

func main() {
	mySlice := []string{"a", "b", "c", "d", "h", "i"}
	fmt.Println(mySlice)
	fmt.Println(mySlice[2:4])  //from 2 to 4 (exclude 4 - position) slicing a slice
	fmt.Println(mySlice[2])    // value by position 2
	fmt.Println("myString"[2]) //index access asci 83 = S
}
