package main

import "go-leanring/oop1"

type Number int

func (n Number) Equal(i int) bool {
	return int(n) == i
}
func (n Number) LessThan(i int) bool {
	return int(n) < i
}

func (n Number) MoreThan(i int) bool {
	return int(n) > i
}

func main() {
	var num1 Number = 1
	var num2 oop1.Number1 = num1
	var num3 oop1.Number2 = num2
}
