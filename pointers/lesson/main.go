package main

import "fmt"

type person struct {
	fname string
	lname string
	age   int
}

func changeMe(p *person) {
	p.fname = "Noyan"
	p.lname = "Alimov"
	p.age = 22
}

func main() {
	pers := person{fname: "John", lname: "Smith", age: 30}
	fmt.Println(pers)

	changeMe(&pers)
	fmt.Println(pers)
}
