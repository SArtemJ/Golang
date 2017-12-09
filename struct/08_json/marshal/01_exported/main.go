package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First        string
	Last         string
	Age          int
	notExpoerted int
}

func main() {
	p1 := Person{"John", "Doe", 20, 001}
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))
}
