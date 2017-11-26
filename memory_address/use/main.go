package main

import "fmt"

const grammTokg float64 = 1000.00

func main() {
	var kgValue float64
	fmt.Println("Enter kg's value: ")
	fmt.Scan(&kgValue)
	result := kgValue * grammTokg
	fmt.Println(result, " gramm in ", kgValue, " kg's")
}
