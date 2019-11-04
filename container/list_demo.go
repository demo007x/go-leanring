package main
import (
	"container/list"
	"fmt"
)

//type List struct {
//	root Element
//	len int
//}
//
//// 链表就是有一个prev和next的指针数组
//type Element struct {
//	next, prev *Element // 上一个和下一个元素
//	list *List
//	value interface{} // 元素
//}

func main() {
	list := list.New()
	// 链表后面插入俩值
	list.PushBack(1)
	list.PushBack(2)

	fmt.Printf("lenL %v\n", list.Len())
	fmt.Printf("first: %#v\n", list.Front())
	fmt.Printf("secibd: %#v\n", list.Front().Next())
}
