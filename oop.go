package main

import (
	"fmt"
	"time"
)

func main() {
	var startAt = "2023-08-01 09:10:11"
	var endAt = "2023-09-02 09:10:11"

	// var startDate = strings.Split(startAt, " ")
	// var endDate = strings.Split(endAt, " ")

	startP, _ := time.ParseInLocation(time.DateTime, startAt, time.Local)
	endP, _ := time.ParseInLocation(time.DateTime, endAt, time.Local)
	for startP.Unix() <= endP.Unix() {
		fmt.Println(startP.Day())
		startP = startP.Local().Add(time.Hour * 24)
	}
}
