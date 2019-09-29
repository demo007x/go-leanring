package main

import "fmt"

// 初始化内嵌结构体
type Wheel struct {
	Size int
}

type Car struct {
	Wheel
	Engine struct{
		Power int
		Type string
	}
}

func main() {
	c := Car{
		Wheel: Wheel{
			Size:18,
		},
		Engine: struct {
			Power int
			Type  string
		}{
			143,
			"1.4T",
		},
	}

	fmt.Printf("%+v \n", c)
}



