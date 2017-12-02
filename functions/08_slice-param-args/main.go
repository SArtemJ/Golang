package main

import "fmt"

func main() {
	data := []float64{43, 56, 87, 12, 45, 57}
	n := average(data)
	fmt.Println(n)
}

func average(sf []float64) float64 {
	var result float64
	for _, value := range sf {
		result += value
	}
	return result / float64(len(sf))
}
