package main

//type Circle struct {
//	X, Y, Radius int
//}
//
//type Wheel struct {
//	X, Y, Radius, Spokes int
//}


type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	//var w Wheel
	//w.X = 8
	//w.Y = 8
	//w.Radius = 5
	//w.Spokes = 20
	
	w = Wheel{
		Circle: Circle{
			Point{X: 10, Y: 2},
			Radius:20,
		},
		Spokes: 20,
	}
}
