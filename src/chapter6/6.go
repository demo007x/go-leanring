package main

import "fmt"

type Point struct {
	X int
	Y int
}

func (p Point) Add(other Point) Point {
	return Point{
		X: p.X + other.X,
		Y: p.Y + other.Y,
	}
}

func main() {
	p1 := Point{1,1}
	p2 := Point{2,2}

	result := p1.Add(p2)

	fmt.Println(result)
}


