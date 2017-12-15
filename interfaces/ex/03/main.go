package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{4, 8, 9, 1, 2, 4, 5, 6, 7, 1}

	fmt.Println(n)
	sort.Sort(sort.IntSlice(n))
	fmt.Println(n)
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(n)
}
