package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIN(boring("John"), boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boting: I'm 1leaving")
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			//запись знач в канал
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Microsecond)
		}
	}()
	return c
}

func fanIN(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			//берем значение из input1 и помещаем его в канал с
			c <- <-input1
		}
	}()
	go func() {
		c <- <-input2
	}()
	return c
}
