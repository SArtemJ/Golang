package main

import (
	"fmt"
	"os"
)

/*func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
*/

/*
func main() {
	s, sep  := "", ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(srings.Join(os.Args[1:], " "))
}*/

func main() {
	s, separator := "", ""

	/*for i := 0; i < len(os.Args); i++ {
		s += separator + os.Args[i]
		separator = "_"
	}*/
	for a, b := range os.Args[0:] {
		fmt.Println(a)
		s += separator + b
		separator = "	"
		fmt.Println(s)
	}


}
