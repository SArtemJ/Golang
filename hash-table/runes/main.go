package main

import "fmt"

func main() {
	//01 letter := 'A' //' ' letter " " - string
	//02
	//letter := rune("A"[0])
	//fmt.Println(letter)
	//01 fmt.Printf("%T \n", letter)

	//03
	word := "Hello"
	letter := rune(word[0])
	fmt.Println(letter)

}
