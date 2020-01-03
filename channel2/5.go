package main

import "fmt"

func main() {
	ch := make(chan int, 11)
	ch <- 99

	for i := 0; i<10; i++ {
		ch<-i
	}
	// 输出ch中的值
	fmt.Printf("writed chan ch : %v\n", ch)
	//	读取chan
	first_chan, ok := <-ch
	if ok {
		fmt.Printf("first chan is %v\n", first_chan)
	}
	ch <- 10
	// 遍历 chan
	for val := range ch {
		fmt.Println(val)
		if val == 10 {
			close(ch)
		}
	}
	fmt.Println("after range or close ch!")
}
