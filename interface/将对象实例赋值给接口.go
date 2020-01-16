package main

import "fmt"

type Integer int

func (a Integer) Equal(i Integer) bool {
	return a == i
}

//func (a *Integer) Equal(i Integer) bool {
//	return (*a).Equal(i)
//}

func (a Integer) LessThan(i Integer) bool {
	return a < i
}

func (a Integer) MoreThan(i Integer) bool {
	return a > i
}

func (a *Integer) Increase(i Integer) {
	*a = *a + i
}

func (a *Integer) Decrease(i Integer) {
	*a = *a - i
}

// 接口
type IntNumber interface {
	Equal(i Integer) bool
	LessThan(i Integer) bool
	MoreThan(i Integer) bool
	Increase(i Integer)
	Decrease(i Integer)
}

type IntNumber2 interface {
	Equal(i Integer) bool
	LessThan(i Integer) bool
	MoreThan(i Integer) bool
}

func main() {
	var a Integer = 1
	var b1 IntNumber = &a
	var b2 IntNumber2 = a
	fmt.Println(a, b1, b2)
}
