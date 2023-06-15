package main

import (
	"encoding/json"
	"fmt"
)

type MyData struct {
	One   string `json:"one"`
	Two   string `json:"two"`
	Three string `json:"three"`
}

func main() {
	in := &MyData{One: "first", Two: "second"}

	var inInterface map[string]string
	inrec, _ := json.Marshal(in)
	json.Unmarshal(inrec, &inInterface)
	fmt.Println(inInterface)
	// iterate through inrecs
	for field, val := range inInterface {
		if val != "" {
			fmt.Println("KV Pair: ", field, val)
		}
	}
}
