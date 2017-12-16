package main

import (
	"fmt"
	"time"
)

func main() {
	//unbuffered channel
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
			//get value and stop go next go-r
		}
	}()

	go func() {
		for {
			fmt.Println(<-c)
			//print and stop
		}
	}()

	time.Sleep(time.Second)
}
