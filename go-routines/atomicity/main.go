package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wait sync.WaitGroup
var counter int64

//var mutex sync.Mutex

func main() {

	wait.Add(2)
	go incrementator("One: -")
	go incrementator("Two: -")
	wait.Wait()
	fmt.Println("Final counter: -", counter)
}

func incrementator(s string) {
	for i := 0; i < 20; i++ {
		//x := counter
		//x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		//mutex.Lock()
		//counter++
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "Counter: ", counter)
		//mutex.Unlock()
	}
	wait.Done()
}

//go run -race main.go
