package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) fullName() string {
	return p.first + p.last
}

func main() {
	p1 := person{"James ", "Ooh", 32}
	p2 := person{"Karla ", "Boom", 27}
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
}
