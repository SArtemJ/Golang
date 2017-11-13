package someMore

import "fmt"

//upperCase to see in files package
//lowerCase can't see inside package
func Max(x int) int {
	return 55 + x

}

func main() {
	//max := max(10) very bad
	result := Max(10)
	fmt.Println(result)
}
