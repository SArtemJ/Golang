package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	/* if try this
		<-done
		<-done
		close(c)
	we can't take value off of channel c
	because we blocked channel
	 */

	go func() {
		<-done
		<-done
		close(c)
	}()



	for n := range c {
		fmt.Println(n)
	}
}
