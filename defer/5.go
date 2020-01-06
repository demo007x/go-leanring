package main

import (
	"fmt"
	"sync"
)

type rect struct {
	length int
	width int
}

func (r rect) area(group *sync.WaitGroup)  {
	defer group.Done()
	if r.length < 0 {
		fmt.Printf("rect %v's length should be greater than zero\n", r)
		return
	}

	if r.width < 0 {
		fmt.Printf("rect %v's width should be greater than zero\n", r)
		return
	}

	area := r.length * r.width
	fmt.Printf("rect %v's area %d",r, area)
	group.Done()
}

func main() {
	var group sync.WaitGroup
	r1 := rect{
		length: -67,
		width:  89,
	}
	r2 := rect{
		length: 5,
		width:  -67,
	}

	r3 := rect{
		length: 8,
		width:  9,
	}

	rects := []rect{r1, r2, r3}
	for _, r :=  range rects {
		group.Add(1)
		go r.area(&group)
	}
	group.Wait()
	fmt.Println("All go routines finished executing")
}