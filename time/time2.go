package main

import (
	"fmt"
	"time"
)

func main() {
	currLocalTime := time.Now()
	curUTCTime := currLocalTime.UTC()
	fmt.Println(currLocalTime)
	fmt.Println(curUTCTime)
	// 设置时区
	fmt.Println(time.Now().In(time.UTC))
	// format
	localTimeStr := time.Now().Format(time.RFC3339)
	utcTimeStr := time.Now().UTC().Format(time.RFC3339)
	fmt.Println(localTimeStr)
	fmt.Println(utcTimeStr)
	// 时间戳
	timestamp := time.Now().Unix()
	timestampNano := time.Now().UnixNano()
	fmt.Println(timestamp)
	fmt.Println(timestampNano)

	// 时间戳转换
	timestamp = time.Now().Unix()
	localTimeObj := time.Unix(timestamp, 0)
	fmt.Println(localTimeObj.Format(time.RFC3339))

	// 时间字符串转换时间类型
	timeStr := "2020-01-13 22:32:17"
	utcTimeObj, _ := time.Parse(time.RFC3339, timeStr)
	fmt.Println(utcTimeObj, utcTimeObj.Unix())

	// 时间+/-
	curTime := time.Now()
	addSecondTime := curTime.Add(time.Second * 1)
	fmt.Println(addSecondTime)
	addMinuteTime := curTime.Add(time.Minute)
	fmt.Println(addMinuteTime)
	addMinuteTime2 := curTime.Add(time.Minute * time.Duration(60))
	fmt.Println(addMinuteTime2)

	// cron
	ticker := time.NewTicker(time.Second * 5) // 5's 后执行
	go func() {
		fmt.Println("println ticker..")
		for _ = range ticker.C {
			fmt.Println("do the something...")
		}
	}()
	fmt.Println("waiting ticker timeout")
}
