package main

import "fmt"

func main() {

	myGreeting := map[int]string{
		0: "Hello",
		1: "Hi",
		2: "Bonjour",
		3: "What's UP",
	}

	fmt.Println(myGreeting)
	//delete(myGreeting, 2)

	if val, exists := myGreeting[2]; exists {
		fmt.Println("val: ", val)
		fmt.Println("existis:", exists)
	} else {
		fmt.Println("That value doesn't exist.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}

	fmt.Println(myGreeting)

}
