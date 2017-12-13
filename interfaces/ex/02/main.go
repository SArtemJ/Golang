package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Qwerty", "Bob", "Bill", "Ulma", "Alis"}
	//fmt.Println(s)
	//sort.Sort(sort.Reverse(sort.StringSlice(s)))
	//fmt.Println(s)

	sort.Sort(sort.StringSlice(s))
	fmt.Println(s)
	fmt.Printf("%T \n", s)
	fmt.Printf("%T \n", sort.StringSlice(s))
	//fmt.Printf("%T \n", sort.Sort(sort.StringSlice(s)))
	fmt.Printf("%T \n", sort.Reverse(sort.StringSlice(s)))
}
