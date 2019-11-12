package main

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	// startA := time.Now()
	// http.Get(a)
	// aDuration := time.Since(startA)

	// startB := time.Now()
	// http.Get(b)
	// bDuration := time.Since(startB)

	// if aDuration < bDuration {
	// 	// a 的时间小于b
	// 	return a
	// }

	// return b
	// 知识点：
	// time.Now() 获取当前的时间
	// http.Get() 请求URL的内容，这个函数返回一个http.Response 和一个error，但我们不关心他们的值
	// time.Since， 传入开始时间并返回一个时间差

	// 简单代码公用：
	// aDuration := measureResponseTime(a)
	// bDuration := measureResponseTime(b)

	// if aDuration < bDuration {
	// 	return a
	// }

	// return b

	// 使用select构造器
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b

	}
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()

	return ch
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
