package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wait sync.WaitGroup

func init() {
	//use all cores
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {

	wait.Add(2)
	go testOne()
	go testTwo()
	wait.Wait()
}
func testOne() {
	for i := 0; i < 35; i++ {
		fmt.Println("Test one - ", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wait.Done()
}

func testTwo() {
	for i := 0; i < 35; i++ {
		fmt.Println("Test two - ", i)
		time.Sleep(time.Duration(20 * time.Millisecond))
	}
	wait.Done()
}
