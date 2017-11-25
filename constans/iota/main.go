package main

import "fmt"

//const (
//
//	//iota start with 0 - increment
//	One   = iota //0
//	//Two   = iota
//	//Three = iota //2
//	Two
//	Three
//)

const (
	_ = iota
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
	GB = 1 << (iota * 10)
)

func main() {

	//fmt.Println(One)
	//fmt.Println(Two)
	//fmt.Println(Three)
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)


}
