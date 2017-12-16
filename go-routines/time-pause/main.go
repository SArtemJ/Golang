package main

import (
	"fmt"
	"sync"
	"time"
)

var wait sync.WaitGroup

func main() {

	wait.Add(2)
	go testOne() //2-1
	go testTwo() //1-1 = 0 wait
	wait.Wait()
}
func testOne() {
	for i := 0; i < 35; i++ {
		fmt.Println("Test one - ", i)
		time.Sleep(time.Duration(3*time.Millisecond))
	}
	wait.Done()
}

func testTwo() {
	for i := 0; i < 35; i++ {
		fmt.Println("Test two - ", i)
		time.Sleep(time.Duration(20*time.Millisecond))
	}
	wait.Done()
}
