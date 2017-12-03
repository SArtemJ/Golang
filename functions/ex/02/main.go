package main

import "fmt"

//func twoReturnValue() func(x int) (int, bool) {
//	return func(x int) (int, bool) {
//		var z1 int
//		var z2 bool
//		if x%2 == 0 {
//			z1 = x / 2
//			z2 = true
//		} else {
//			z1 = x / 2
//			z2 = false
//		}
//		return z1, z2
//	}
//}

func main() {
	x := func(x int) (int, bool) {
		var z1 int
		var z2 bool
		if x%2 == 0 {
			z1 = x / 2
			z2 = true
		} else {
			z1 = x / 2
			z2 = false
		}
		return z1, z2
	}
	fmt.Println(x(10))
}
