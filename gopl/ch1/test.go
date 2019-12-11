package main

import "fmt"

func main() {
	x := 1
	p := &x
	incr(p)
	fmt.Println(incr(&x))
}

func incr(p *int) int {
	*p++
	return *p
}
