package main

import "fmt"

func main() {

	m := make(map[string]int)
	m["key1"] = 100
	m["key2"] = 101
	fmt.Println(m)

	n := map[int]int{10:100, 20:200, 30:300}
	fmt.Println(n)

	delete(m, "key2")
	fmt.Println("my map:", m)

	value, key := m["key1"]
	fmt.Println(key, value)
}
