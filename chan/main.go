package main
import "fmt"
func main()  {
	// c := make(chan int, 1)
	// c <- 10 // 写入数据到c中
	// v := <- c // 读取c中的值给v

	// fmt.Println(v)
	// 判断chan是否关闭：
	// c := make(chan int, 1)
	// c <- 10
	// v, ok := <- c
	// close(c)
	// vv, ok := <- c

	// fmt.Println(v, vv, ok) // 10 0 false
	// 如果写不进去就丢弃， 可以使用 select
	c := make(chan int, 1)
	select {
	case c <- 10:
		fmt.Printf("%#v \n", c)
	default:
	}

	select {
	case c <- 11:
		fmt.Printf("%#v \n", c)
	default: // c 中只有10， 没有11
	}

	select {
	case v, ok := <- c:
		// 读出一个 v = 10, ok = true
		fmt.Printf("%v %t\n", v, ok)
	default:
	}

	select {
	case v, ok := <- c:
		fmt.Printf("%v %t\n", v, ok)
	default:
	}
}