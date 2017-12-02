package main

import "fmt"

func main() {
	//s := greet("Iron ", "Man")
	//fmt.Println(greet("Iron ", "Man"))
	fmt.Print(greet("Iron ", "Man"))
}

func greet(fname, lname string) (string, string) {
	//nameReturnValue = fmt.Sprint(fname, lname)
	return fmt.Sprint(fname, lname), fmt.Sprint(fname, lname)
}
