package main

import (
	"fmt"
)

// 单项channel

func main01() {
	c := make(chan int, 3)
	var send chan <- int = c // send-only, 只能给chan中发送数据
	var recv <-chan int = c  // receive-only 只能接受chan中的数据

	send <- 2
	val, ok := <-recv
	if ok {
		fmt.Println(val)
	}
}

func main02() {
	// 不能将单项chan转换为普通的chan
	//c := make(chan int, 3)
	//var send chan<- int = c
	//var recv <-chan int = c

	//cannot convert send (type chan<- int) to type chan int
	//ch1 := (chan int)(send)
	// cannot convert recv (type <-chan int) to type chan int
	//ch2 := (chan int)(recv)
}

type Request struct {
	data []int
	ret chan int
}

func NewRequest(data ...int) *Request {
	return &Request{
		data: data,
		ret:  make(chan int, 1),
	}
}

func Process(req *Request)  {
	x := 0
	for _, i := range req.data {
		x += i
	}
	req.ret <- x
}

func main() {
	req := NewRequest(10, 20, 30)
	Process(req)
	fmt.Println(<-req.ret)
}

