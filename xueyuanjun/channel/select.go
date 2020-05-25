package main

import (
	"fmt"
	"math/rand"
)

func main() {
	chs := [3]chan int {
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}

	index := rand.Intn(3) //生成随机的0-2的值
	fmt.Printf("随机值 %d \n", index)
	chs[index]<-index// 写入通道的值

	select {
	case <- chs[0]:
		fmt.Println("第一个分支被选中。。。")
	case <- chs[1]:
		fmt.Println("第二个分支被选择...")
	case num := <- chs[2]:
		fmt.Println("第三个分支被选中：", num)
	default:
		fmt.Println("默认分支执行")
	}
}