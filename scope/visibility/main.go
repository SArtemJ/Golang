package main

import ("fmt"
	"github.com/SArtemJ/GoLang/scope/visibility/someMore"
)

func main() {
	x := 55
	fmt.Println(x)
	fmt.Println(y)
	secondFunc()
}

func secondFunc() {
	//fmt.Println(x)
	fmt.Println(someMore.max(11))
}

var y = 101
