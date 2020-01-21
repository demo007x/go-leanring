package main

import (
	"fmt"
	"time"
)

func main01() {
	timer1 := time.NewTimer(time.Second * 2) // 到时间才会执行的一个通道
	<-timer1.C
	fmt.Println("Timer 1 expired")
}

func main02() {
	timer2 := time.NewTimer(time.Second * 2)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()

	stop := timer2.Stop()

	if stop {
		fmt.Println("Timer 2 stopped")
	}
}

func main03() {
	ticker := time.NewTicker(time.Microsecond * 500)
	go func() {
		fmt.Println("go func...")
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Second * 1)
}

func main04() {
	go func() {
		time.Sleep(time.Hour)
	}()
	c := make(chan int, 10)
	c <- 1
	c <- 2

	close(c)
	c <- 3
}

func worker(ch chan int) {
	time.Sleep(time.Second * 2)
	ch <- 3
}

func main() {
	ch := make(chan int, 1)
	go worker(ch)
	// 等待任务完成
	r := <-ch
	fmt.Println(r)

}
