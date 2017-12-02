package main

import "fmt"

func main() {
	s := greet("Iron ", "Man")
	//fmt.Println(greet("Iron ", "Man"))
	fmt.Print(s)
}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}
