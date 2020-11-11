package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Fname string `json:"first name"`
	Lname string `json:"last name"`
	Age   int    `json:"age"`
}

func main() {
	p := person{"John", "Doe", 31}
	bs, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

	j := `{"first name":"John","last name":"Doe","age":31}`
	bsl := []byte(j)
	var p2 person
	err = json.Unmarshal(bsl, &p2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p2)
}
