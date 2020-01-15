package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Unix())
	fmt.Println(t.String())

	// time.Duration
	fmt.Println(time.Duration(100000) * time.Second)

	// timer
	fmt.Println(time.Now())
	timer := time.NewTicker(time.Second * 2) // 2秒收到期
	<-timer.C
	fmt.Println(time.Now())
}
