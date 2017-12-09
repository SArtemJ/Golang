package main

import "fmt"

type Person struct {
	First string
	Last  string
	Age   int
}

type DoubleZero struct {
	Person
	DriveLicense bool
}

func main() {
	p1 := DoubleZero{
		Person: Person{
			First: "Bob",
			Last:  "Crypto",
			Age:   27,
		},
		DriveLicense: true,
	}

	p2 := DoubleZero{
		Person: Person{
			First: "Klara",
			Last:  "Boom",
			Age:   25,
		},
		DriveLicense: false,
	}

	//p3 := DoubleZero{
	//	_: Person{
	//		First: _,
	//		Last:  _,
	//		Age:   _,
	//	},
	//	DriveLicense: false,
	//}

	fmt.Println(p1.First, p1.Last, p1.Age, p1.DriveLicense)
	fmt.Println(p2.First, p2.Last, p2.Age, p2.DriveLicense)
	//fmt.Println(p3.First, p3.Last, p3.Age, p3.DriveLicense)
}
