package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int {
	return len(p)

}

func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]

}

func (p people) Less(i, j int) bool {
	return p[i]<p[j]
}

func main() {
	k := people{ "Zdash", "Alis", "Bob", "Klara", "James"}
	fmt.Println(k)
	sort.Sort(k)
	fmt.Println(k)

	fmt.Println()

	s := []string{"James", "Alis", "Bob", "Klara",  "Brash"}
	fmt.Println(s)
	//sort.Sort(people(s))
	//sort.StringSlice(s).Sort()
	//sort.Sort(sort.StringSlice(s))
	sort.Strings(s)
	fmt.Println(s)


}
