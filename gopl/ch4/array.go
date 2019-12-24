package main

import "fmt"

// var a [3]int = [3]int{1, 2, 3}
// var r [3]int = [3]int{1, 2}
// var q = [...]int{1, 2, 3, 4}

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {

	// fmt.Println(a[0])
	// fmt.Println(a[len(a)-1])
	// fmt.Println(r[2])

	// for _, v := range q {
	// 	fmt.Println(v)
	// }

	// symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}

	// fmt.Println(RMB, symbol[RMB])

	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c)
	d := [3]int{1, 2}
	// fmt.Println(a == d) // false
}
