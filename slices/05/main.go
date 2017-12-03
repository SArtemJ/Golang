package main

import "fmt"

func main() {

	sliceInts := make([][]int, 0, 3)

	for i := 0; i < 3; i++ {
		sliceInt := make([]int, 0)
		for j := 0; j < 4; j++ {
			sliceInt = append(sliceInt, j)
		}
		sliceInts = append(sliceInts, sliceInt)
	}
	fmt.Println(sliceInts)

}
