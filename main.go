package main

import (
	"fmt"
	"encoding/json"
)

type FizzBuzz struct {
	ValueStr string `json:"value,omitempty"` //тэг value и пишем value
	ValueInt int    `json:"value,omitempty,string"`
}

func main() {
	values := []string{}  //по сути не используется нигде
	for i := 1; i <= 20; i++ {
		v := FizzBuzz{}
		switch {
		case i%15==0:
			v.ValueStr = "FizzBuzz"
		case i%3==0:
			v.ValueStr = "Fizz"
		case i%5==0:
			v.ValueStr = "Buzz"
		default:
			v.ValueInt = i
		}
		b, _ := json.Marshal(v)
		values = append(values, string(b)) //по сути не используется нигде
		fmt.Println(string(b))
	}

	jsonStr := `{"value": 42}` //value и тэг value при json запутается так как sting/int тэга нет
	//должно быть "42"
		fb := FizzBuzz{}
	   json.Unmarshal([]byte(jsonStr), fb)  //тут надо передавать по &
		fmt.Println(fb)
}
