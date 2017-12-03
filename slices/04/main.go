package main

import "fmt"

func main() {
	journal := make([][]string, 0)
	clientOne := make([]string, 4)
	clientOne[0] = "John"
	clientOne[1] = "Doe"
	clientOne[2] = "111-222-333"
	clientOne[3] = "Last Week"

	journal = append(journal, clientOne)

	clientTwo := make([]string, 4)
	clientTwo[0] = "Elliot"
	clientTwo[1] = "Alderson"
	clientTwo[2] = "444-555-666"
	clientTwo[3] = "Mr. Robot"

	journal = append(journal, clientTwo)
	fmt.Println(journal)
}
