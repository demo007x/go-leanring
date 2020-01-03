package main

import "fmt"

type Stu struct {
	name string
}

func main() {
	var intChan chan int
	intChan = make(chan int, 10)
	intChan<-10

	a := <-intChan
	fmt.Printf("int 类型chan : %v\n", a)

	// map 类型
	var mapChan chan map[string]string
	mapChan = make(chan map[string]string, 10)
	m := make(map[string]string, 16)
	m["stu01"] = "001"
	m["stu02"] = "002"
	m["stu03"] = "003"
	mapChan<-m
	b := <-mapChan
	fmt.Printf("map 类型chan ： %v\n", b)
	// 结构体
	var stuChan chan Stu
	stuChan = make(chan Stu, 10)
	stu := Stu{name:"test"}
	stuChan<-stu
	tempStu :=  <- stuChan
	fmt.Printf("struct类型的chan : %v\n", tempStu)
	// 结构体内存地址
	var stuChanId chan *Stu
	stuChanId = make(chan *Stu, 10)
	stuId := &Stu{name:"asdfadsf"}
	stuChanId <- stuId
	tempStuId :=  <- stuChanId
	fmt.Printf("*struct 类型 chan: %v\n", tempStuId)
	fmt.Printf("*struct 类型 chan 取值: %v\n", *(tempStuId))
	// 接口
	var StuInterChan chan interface{}
	StuInterChan = make(chan interface{}, 10)
	stuInit := Stu{name:"63456465"}
	// 存
	StuInterChan <- &stuInit
	// 取
	mFetchStu := <-StuInterChan
	fmt.Printf("interface 类型 chan : %v\n", mFetchStu)

	// 转
	var mStuConvert *Stu
	mStuConvert, ok := mFetchStu.(*Stu)
	if !ok {
		fmt.Println("cannot convert")
		return
	}
	fmt.Printf("interface chan转 *struct chan : %v\n", mStuConvert)
	fmt.Printf("interface chan转 *struct chan 取值 : %v\n", *(mStuConvert))
}
