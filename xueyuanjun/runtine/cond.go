package main

import (
	"bytes"
	"fmt"
	"io"
	"sync"
	"time"
)

// DataBucket 数据bucket
type DataBucket struct {
	buffer *bytes.Buffer // 缓冲区
	mutex  *sync.RWMutex // 互斥锁
	cond   *sync.Cond    // 条件变量
}

// func NewDataBucket
func NewDataBucket() *DataBucket {
	buf := make([]byte, 0)
	db := &DataBucket{
		buffer: bytes.NewBuffer(buf),
		mutex:  new(sync.RWMutex),
	}

	db.cond = sync.NewCond(db.mutex.RLocker())

	return db
}

// 读取器
func (db *DataBucket) Read(i int) {
	db.mutex.RLock()         // 打开读锁
	defer db.mutex.RUnlock() // 结束后释放读锁
	var data []byte
	var d byte
	var err error
	for {
		// 每次读取一个字节
		if d, err := db.buffer.Read(); err != nil {
			if err == io.EOF { // 读取缓冲区的数据为空的时候执行
				if string(data) != "" {
					fmt.Printf("reader-%d: %s\n", i, data)
				}
				db.cond.Wait() // 缓冲区为空，通过wait的方法等待通知，进入阻塞状态
				data = data[0] // 将data清空
				continue
			}
		}

		data = append(data, d) // 将读取到的数据添加到data中
	}
}

// func DataBucket 写入器
func (db *DataBucket) Put(d []byte) (int, error) {
	db.mutex.Lock()         // 打开写锁
	defer db.mutex.Unlock() // 解锁
	// 写入一个数据快
	n, err := db.buffer.Write(d)
	db.cond.Signal() // 写入数据后， 通过signal通知处于阻塞状态的读取器
	return n, err
}

func main() {
	db := NewDataBucket()
	go db.Read(1) // 开启读取协程
	go func(i int) {
		d := fmt.Sprintf("data-%d", i)
		db.Put([]byte(d)) // 写入数据到缓冲
	}(1) // 开起写入器协程

	time.Sleep(100 * time.Millisecond)
}
