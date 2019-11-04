package main

import (
	"errors"
	"os"
	"sync"
)

var (
	valueByKey      = make(map[string]int)
	valueByKeyGuard sync.Mutex
)

//func readValue(key string) int  {
//	valueByKeyGuard.Lock()
//
//	v := valueByKey[key]
//	// 对共享资源解锁
//	valueByKeyGuard.Unlock()
//	return v
//}

func readValue(key string) int {
	valueByKeyGuard.Lock()

	defer valueByKeyGuard.Unlock()
	return valueByKey[key]
}

//func fileSize(filename string) int64  {
//	f, err := os.Open(filename)
//
//	if err != nil {
//		return 0
//	}
//	// 获取文件的状态信息
//	info, err := f.Stat()
//	if err != nil {
//		f.Close()
//		return  0
//	}
//
//	size := info.Size()
//	f.Close()
//	return size
//}

func fileSize(filename string) int64 {
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return 0
	}

	info, err := f.Stat()

	if err != nil {
		return 0
	}

	size := info.Size()

	return size
}

var err = errors.New("this is an error")
