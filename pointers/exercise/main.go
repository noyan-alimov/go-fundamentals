package main

import "fmt"

type person struct {
	name string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Speak")
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{"John"}
	saySomething(&p1)
	fmt.Printf("%T\n", &p1)
}
