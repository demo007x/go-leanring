package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

const HOST =  "localhost:6379"

//func main1 (){
//	c, err := redis.Dial("tcp", "localhost:6379")
//	if err != nil {
//		fmt.Println("conn redis failed,", err)
//		return
//	}
//
//	fmt.Println("redis conn success")
//
//	defer c.Close()
//
//	_, err = c.Do("Set", "abc", 100)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	r, err := c.Do("Get", "abc")
//	if err != nil {
//		fmt.Println("get abc failed",err)
//		return
//	}
//
//	fmt.Println(r)
//}

func main02()  {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed", err)
		return
	}

	defer c.Close()
	// 批量操作
	_, err = c.Do("MSet", "abc", 100, "efg", 300)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.Ints(c.Do("MGet", "abc", "efg"))
	if err != nil {
		fmt.Println("get abc failed", err)
		return
	}

	for _, v := range r {
		fmt.Println(v)
	}
}

// 设置过期时间
func main03() {
	c, err := redis.Dial("tcp", HOST)
	if err != nil {
		fmt.Println("conn redis failed", err)
		return
	}
	defer c.Close()
	_, err = c.Do("Set", "abc", 1000)
	if err != nil {
		fmt.Println(err)
		return;
	}
	_, err = c.Do("expire", "abc", 10)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main04()  {
	c, err := conn()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer c.Close()
	_, err = c.Do("lpush", "book_list", "abc", "ceg", 300)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.String(c.Do("lpop", "book_list"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}

	fmt.Println(r)
}

func conn() (redis.Conn, error) {
	c, err := redis.Dial("tcp", HOST)

	return c, err
}

func main()  {
	conn, err := conn()
	if err != nil {
		fmt.Println("conn redis failed, ",err)
		return
	}
	defer conn.Close()
	_, err = conn.Do("HSet", "books", "abc", 100)

	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.Int(conn.Do("HGet", "books", "abc"))
	if err != nil {
		fmt.Println("get abc failed, ", err)
		return
	}

	fmt.Println(r)

	hAll, err := conn.Do("HGetAll", "books")

	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := redis.ByteSlices(hAll, err)

	for _, v := range result {
		fmt.Println(redis.String(v, err))
	}

}
