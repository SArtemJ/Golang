package main

import "fmt"

type Person struct {
	First string
	Last  string
	Age   int
}

type DoubleZero struct {
	Person
	First        string
	DriveLicense bool
}

func (p Person) Greeting() {
	fmt.Println("Hello person p")
}

func (dz DoubleZero) Greeting() {
	fmt.Println("Hello dz")
}

func main() {
	p1 := DoubleZero{
		Person: Person{
			First: "Bob",
			Last:  "Crypto",
			Age:   27,
		},
		First:        "Double zero first",
		DriveLicense: true,
	}

	p2 := DoubleZero{
		Person: Person{
			First: "Klara",
			Last:  "Boom",
			Age:   25,
		},
		First:        "Double zero first",
		DriveLicense: false,
	}

	p1.Person.Greeting()
	p2.Greeting()

}
