package main

import "fmt"

type Wheel struct {
	Size int
}

// 引擎
type Engine struct {
	Power int
	Type string
}

type Car struct {
	Wheel
	Engine
}

func main() {
	c := Car{
		Wheel:  Wheel{
			Size:18,
		},
		Engine: Engine{
			Power:143,
			Type:"1.4T",
		},
	}

	fmt.Printf("%+v \n", c)
}

