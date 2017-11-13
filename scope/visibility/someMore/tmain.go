package someMore

import "fmt"

//upperCase to see in this package
//lowerCase can't see outside this package
func Max(x int) int {
	return 55 + x

}

func main() {
	//max := max(10) very bad
	result := Max(10)
	fmt.Println(result)
}
