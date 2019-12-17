package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells = 100
	xyrange = 30.0
	xyscale = width / 2 / xyrange
	angle = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) //sin(30), cos(30)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

}