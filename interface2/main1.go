package main

import "fmt"

type A interface {
	Foo()
}

type B interface {
	A
	Bar()
}

type T struct {
}

func (t T) Foo() {
	fmt.Println("Call Foo function form interface A")
}

func (t T) Bar() {
	fmt.Println("Call Bar function form interface B")
}

func main() {
	var t = T{}
	var a A = t
	a.Foo()

	var b B = t
	b.Foo()
	b.Bar()
}
