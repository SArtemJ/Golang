package main

import "fmt"

func main() {

	var myGreeting = make(map[string]string)
	//if -> var myGreeting = map[string]string
	//than -> myGreeting = nil
	myGreeting["Bob"] = "Hello"
	myGreeting["Alis"] = "Aloha"

	fmt.Println(myGreeting)
}
