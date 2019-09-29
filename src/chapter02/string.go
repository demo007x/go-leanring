package main

import (
	"fmt"
	"math"
)

func main() {
	//tracer := "死神来了， 死神bye bye"
	//comma := strings.Index(tracer, "，")
	//fmt.Printf("Index of %d %T \n", comma, comma)
	//fmt.Printf("%s \n", tracer[comma:])
	//pos := strings.Index(tracer[comma:], "死神")
	//fmt.Printf("pos value is %d \n", pos)
	//fmt.Println(comma, pos, tracer[comma+pos:])

	//angel := "Heros never die"
	//angleBytes := []byte(angel)
	//fmt.Printf("%T \n", angleBytes)
	//for i := 5; i <= 10; i++ {
	//	angleBytes[i] = ' '
	//}
	//
	//fmt.Println(string(angleBytes))

	// 字符串拼接
	//hammer := "吃我一锤,"
	//sickle := "受死吧"
	//var stringBuilder bytes.Buffer
	//
	//stringBuilder.Write([]byte(hammer))
	//stringBuilder.Write([]byte(sickle))
	//
	//fmt.Println(stringBuilder.String())

	// 字符串格式化
	var progress  = 2
	var target = 8

	title := fmt.Sprintf("已采集%d个草药， 还需要%d个任务", progress, target)

	fmt.Println(title)

	pi := math.Pi
	// 按数值本身的格式输出
	variant := fmt.Sprintf("%v %v %v", "月球基地", pi, true)
	fmt.Println(variant)

	prifile := &struct {
		Name string
		HP int
	}{
		Name: "rat",
		HP: 150,
	}
	fmt.Printf("使用'%%+v' %+v\n", prifile)
	fmt.Printf("使用'%%#v' %#v\n", prifile)
	fmt.Printf("使用'%%T' %T\n", prifile)
}

