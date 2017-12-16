package main

import (
	"fmt"
	"sync"
)

var wait sync.WaitGroup

func main() {

	//two items to wait
	wait.Add(2)
	//go routine 1 - main
	//go routine 2 - testOne
	go testOne() //2-1
	//go routine 3 - testTwo
	go testTwo() //1-1 = 0 wait
	wait.Wait()
}
func testOne() {
	for i := 0; i < 35; i++ {
		fmt.Println("Test one - ", i)
	}
	wait.Done()
}

func testTwo() {
	for i := 0; i < 35; i++ {
		fmt.Println("Test two - ", i)
	}
	wait.Done()
}
