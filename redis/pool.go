package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

const HOST = "localhost:6379"

var pool *redis.Pool // 创建redis的连接池

func init() {
	// 初始化一个连接池
	pool = &redis.Pool{
		MaxIdle: 16, //最初的一个连接数
		//MaxActive: 10000, // 最大连接数
		MaxActive: 0, // 最大连接数不受控制, 设置为0 按需分配
		IdleTimeout: 3000, // 如果一个链接在300秒内没链接, 则自动关闭
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", HOST)
		},
	}
}

func main() {
	c := pool.Get() // 从连接池里面获取一个链接
	defer c.Close() // 函数运行结束, 把连接池放回到连接池
	_, err := c.Do("Set", "abc", 200)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.Int(c.Do("Get", "abc"))
	if err != nil {
		fmt.Println("get abc failed: ", err)
		return
	}

	fmt.Println(r)
	pool.Close()
}
