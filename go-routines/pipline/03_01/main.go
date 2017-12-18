package main

//ex
import "fmt"

func main() {
	a := getValue()
	b := factorial(a)
	for n := range b {
		fmt.Println(n)
	}
}

func factorial(t <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range t {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func getValue() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i>0; i-- {
		total *= i
	}
	return total
}
