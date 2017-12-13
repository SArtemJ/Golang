package main

import "fmt"

func main() {
	var s = map[int]string{1: `q`, 2: `r`, 3: `t`}
	for _, value := range s {
		fmt.Printf("%c \n", value)
	}
}
