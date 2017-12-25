package main

import (
	"os"
	"fmt"
	//"log"
)

func main()  {
	_, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println("Error from open file")
		//log.Println("Error from open file", err)
		//log.Fatalln(err)
		//panic(err)
	}
}
