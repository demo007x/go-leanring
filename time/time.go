package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	// 有问题
	fmt.Println(time.Parse("2016-01-02 15:04:05", "2018-04-23 12:24:51"))
	fmt.Println(time.Local)
	fmt.Println(time.ParseInLocation("2006-01-02 15:04:05", "2017-05-11 14:06:06", time.Local))
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Unix(time.Now().Unix(), 0))
	fmt.Println(time.Date(2020, 1, 14, 17, 40, 0, 0, time.Local))
	loc, _ := time.LoadLocation("PRC")
	fmt.Println(time.Now().In(loc))

	// to string

	fmt.Println(time.Now().Format(time.RFC3339Nano))
	fmt.Println(time.Now().Format(time.UnixDate))

	// to time stamp
	dt, _ := time.Parse(time.RFC3339, "2018-04-23 12:24:51")
}
