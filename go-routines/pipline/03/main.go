package main

//ex
import "fmt"

func main() {
	a := getValue(5, 6)
	b := factorial(a)
	for n := range b {
		fmt.Println(n)
	}
}

func factorial(n chan int) chan int {
	out := make(chan int)
	go func() {
		for value := range n {
			total := 1
			for i := value; i > 0; i-- {
				total *= i
			}
			out <- total
		}
		close(out)

	}()
	return out
}

func getValue(t ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range t {
			out <- n
		}
		close(out)
	}()
	return out
}
