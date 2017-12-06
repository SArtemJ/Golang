package main

import "fmt"

type Person struct {
	first string
	last  string
	age   int
}

type DoubleZero struct {
	Person
	First          string
	LicenceToDrive bool
}

func main() {
	p1 := DoubleZero{
		Person: Person{
			first: "Steve",
			last:  "Mccuin",
			age:   21,
		},
		First:          "Double Zero Seven",
		LicenceToDrive: true,
	}
	fmt.Println(p1)
}
