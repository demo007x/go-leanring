package main

import (
	"container/ring"
	"fmt"
)

// 还的结构定义
//type Rang struct {
//	next, prev *ring.Ring
//	Value interface{}
//}

func main() {
	// 创建一个长度为3的环结构
	ring := ring.New(3)

	for i := 1; i <= 3; i++ {
		ring.Value = i
		ring = ring.Next()
	}

	// 计算1+2+3
	s := 0
	ring.Do(func(i interface{}) {
		s += i.(int)
	})

	fmt.Println("sum is", s)
}

