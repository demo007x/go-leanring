package main

import (
	"container/list"
	"fmt"
	"sync"
)

var l = list.New()

func main() {
	var scene sync.Map
	// 将值保存在sync.Map中
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)

	// 从 sync.Map 中取值
	fmt.Println(scene.Load("london"))
	// 根据键删除对应的值
	scene.Delete("london")

	scene.Range(func(key, value interface{}) bool {
		fmt.Println("iterate:", key, value)
		return true
	})

	l.PushFront("fist")
	l.PushBack("76")
}
