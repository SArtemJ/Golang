package main

import (
	"fmt"
	"io/ioutil"
	//"log"
	"net/http"
)
//with check error
//func main() {
//
//	res, err := http.Get("http://www.google.com")
//	if err != nil {
//		log.Fatal(err)
//	}
//	page, err := ioutil.ReadAll(res.Body)
//	res.Body.Close()
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("%s", page)
//
//}

//no check error
func main()  {
	res, _ := http.Get("http://www.google.com")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}
