package main

import "fmt"

type persion struct {
	firstName string
	lastName string
}

func (p persion) fullName() {
	fmt.Printf("%s %s", p.firstName, p.lastName)
}

func main() {
	p := persion{
		firstName: "x",
		lastName:  "x",
	}

	defer p.fullName()
	fmt.Printf("welcome \t")
}
