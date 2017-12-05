package main

import "fmt"

func main() {
	n := HashBucket("Go", 12)
	fmt.Println(n)
}

func HashBucket(word string, buckets int) int {
	//letter := rune(word[0]) - need change bucket by int32 and return int32 type
	letter := int(word[0])
	bucket := letter % buckets
	return bucket

}
