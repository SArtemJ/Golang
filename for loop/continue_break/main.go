package main

import "fmt"

func main() {

	i := 0
	//for {
	//	fmt.Println(i)
	//	if i >= 10 {
	//		break
	//	}
	//	i++
	//}

	for {
		i++
		for i%2 == 0 {
			continue
		}
		fmt.Println(i)
		if i >= 50 {
			break
		}
	}

}
