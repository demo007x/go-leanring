package main

import "fmt"

func main() {
	//缓冲信道的长度是指信道中当前排队的元素个数。
	//ch := make(chan string, 2) // 2 是指通道中元素的个数
	//ch <- "naveen"
	//ch <- "paul"
	//ch <- "steve" //fatal error: all goroutines are asleep - deadlock!
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	//close(ch)

	ch := make(chan string, 3)
	ch <- "naveen"
	ch <- "paul"
	fmt.Println("capacity is", cap(ch))
	fmt.Println("length is", len(ch))
	fmt.Println("read value", <-ch)
	fmt.Println("read value", <-ch)
	fmt.Println("new length is ", len(ch))
}
