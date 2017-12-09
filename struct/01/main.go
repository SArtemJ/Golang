package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {

	var x person //zero value fields
	fmt.Println(x)
	p1 := person{"John", "Doe", 30}
	p2 := person{"Elliot", "Alderson", 25}
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
}


