package someMore

import "fmt"


func max(x int) int {
	return 55 + x

}

func main() {
	//max := max(10) very bad
	result := max(10)
	fmt.Println(result)
}
