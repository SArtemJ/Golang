package main

import "fmt"

func main() {

	myGreeting := make(map[string]string)
	myGreeting["Bob"] = "Hello"
	myGreeting["Alis"] = "Aloha"

	//or
	//myGreeting := map[string]string {
	//	"David": "Good day",
	//	"Samanta": "Bonjour!",
	//}

	myGreeting["Leon"] = "What's UP!"
	fmt.Println(myGreeting)
	fmt.Println(myGreeting["Leon"])
}
