package main

import "oop/employee"

func main() {
	//e := employee.Employee{
	//	FirstName:    "anziguoer",
	//	LastName:     "杨玉龙",
	//	TotalLeavers: 30,
	//	LeaversTaken: 20,
	//}
	//
	//e.LeavesRemaining()
	// 定义一个0值的结构体

	e := employee.New("Sam", "Adolf", 30, 20)
	e.LeavesRemaining()
}
