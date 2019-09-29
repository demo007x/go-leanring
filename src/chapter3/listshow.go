package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	// 尾部添加
	l.PushBack("canon")
	// 头部添加
	l.PushFront("67")
	// 尾部添加元素后保存元素的具柄
	element := l.PushBack("first")
	// 在first添加之后天花high
	l.InsertAfter("high", element)
	// 在first添加之前添加 noon
	l.InsertBefore("noon", element)
	// 使用
	l.Remove(element)

	for i := l.Front(); i != nil ; i = i.Next() {
		fmt.Println(i.Value)
	}
}
