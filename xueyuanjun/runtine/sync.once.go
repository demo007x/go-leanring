package main

import (
	"fmt"
	"sync"
	"time"
)

func dosomething(o *sync.Once) {
	fmt.Println("start:")
	o.Do(func() {
		// v只执行一次
		fmt.Println("do something...")
	})

	fmt.Println("finished...")
}

func main() {
	o := &sync.Once{}
	go dosomething(o)
	go dosomething(o)
	time.Sleep(time.Second * 1)
}
