package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i:=0; i<10; i++ {
			c <- i
		}
		close(c)
	}()

	//error dead lock, because we put all values in channel but not close
	//for {
	//	fmt.Println(<-c)
	//}

	//only 0 print because we need close channel and take values with range
	//fmt.Println(<-c)

	for n:=range c {
		fmt.Println(n)
	}

}
