package main

import (
	"fmt"
	"math/rand"
	"time"
)

const formatDate = "2006-01-02 15:04:05"

func main001() {
	t, _ := time.ParseInLocation(formatDate, "2020-02-26 09:14:00", time.Local)
	fmt.Println(time.Local)
	fmt.Println(t.Location())
	fmt.Println(time.Now().Sub(t).Hours())

	// 输出正点的时间
	t,_ = time.ParseInLocation(formatDate, time.Now().Format("2006-01-02 15:04:05"), time.Local)
	fmt.Println(t)


	t, _ = time.ParseInLocation("2006-01-02 15:04:05", "2016-06-13 15:34:39", time.Local)
	// 整点（向下取整）
	fmt.Println(t.Truncate(1 * time.Hour))
	fmt.Println("整分", t.Truncate(1 * time.Minute))
}

func main002() {
	c := make(chan int) // 创建一个通道
	go func() {
		time.Sleep(3 * time.Second)
		<-c // 通道输出
	}()

	select {
	case c <- 1:
		fmt.Println("channel...")
	case <-time.After(2 * time.Second):
		close(c)
		fmt.Println("timeout...")
	}
}

func main003() {
	start := time.Now()
	timer := time.AfterFunc(2 * time.Second, func() {
		// 2 s 后开始执行
		fmt.Println("after func callback, elaspe", time.Now().Sub(start))
	})
	//time.Sleep(2 * time.Second)
	//fmt.Println(timer.Stop())

	//time.Sleep(1 * time.Second)
	//
	if timer.Reset(3 * time.Second) {
		fmt.Println("timer has not trigger!")
	} else {
		fmt.Println("timer had expired or stop!")
	}

	time.Sleep(10 * time.Second)
}

func main004() {
	fmt.Println("当前时间为：", time.Now().Format(formatDate))
	timer := time.AfterFunc(2 * time.Second, func() {
		fmt.Println("time.AfterFunc 2's 后输出的：",time.Now().Format(formatDate))
	})

	time.Sleep(3 * time.Second)
	fmt.Println("3 * time.Second：", time.Now().Format(formatDate))
	fmt.Println("stop :", timer.Stop())
	fmt.Println("stop :, timer.Stop():", time.Now().Format(formatDate))
	fmt.Println(timer.Reset(3 * time.Second))
	fmt.Println("timer.Reset(3 * time.Second):", time.Now().Format(formatDate))

	time.Sleep(10)
}

func main005()  {
	ticker := time.NewTicker(500 * time.Millisecond) // 500 毫秒
	done := make(chan bool)
	
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done<-true
	fmt.Println("Ticker stopped")
}

func main006() {
	d := time.Duration(time.Second * 2)

	t := time.NewTicker(d)
	defer t.Stop()

	for {
		<-t.C
		fmt.Println("timeout...")
	}
}

func main007() {
	ticker := time.NewTicker(time.Second * 1)
	i := 0
	go func() {
		for {
			<-ticker.C
			i++
			fmt.Println("i=", i)
			if i == 5 {
				ticker.Stop()
			}
		}
	}()

	time.Sleep(10 * time.Second)
}

func main008() {
	// 初始化定时器
	t := time.NewTimer(10 * time.Second)
	// 当前时间
	now := time.Now()
	fmt.Printf("Now time : %v \n", now)
	expire := <-t.C
	fmt.Printf("Expiration time: %v \n", expire)
}

func main009() {
	ch11 := make(chan int, 1000)
	sign := make(chan byte, 1)

	//  给 ch11 通道中添加数据
	for i := 0; i < 1000; i++ {
		ch11 <- i
		fmt.Println("i = ", i)
	}

	go func() {
		var e int
		ok := true
		//首先声明一个*time.Timer类型的值，然后在相关case之后声明的匿名函数中尽可能的复用它
		var timer *time.Timer
		for {
			select {
			case e = <- ch11:
				fmt.Printf("ch11 -> %d\n", e)
			case <-func() <-chan time.Time {
				if timer == nil {
					timer = time.NewTimer(time.Millisecond)
				} else {
					timer.Reset(time.Millisecond)
				}

				return timer.C
			}():
				fmt.Println("Timeout...")
				ok = false
				break
			}
			if !ok {
				sign <- 0
				break
			}
		}
	}()
	<-sign
}

func main010() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	rand.Seed(time.Now().Unix())
	n := rand.Intn(3)
	if n == 1 {
		ch1 <- 1
	} else if n == 2 {
		ch2 <- 1
	}
	fmt.Println(n)
	select {
	case e1 := <-ch1:
		//如果ch1通道成功读取数据，则执行该case处理语句
		fmt.Printf("1th case is selected. e1=%v",e1)
	case e2 := <-ch2:
		//如果ch2通道成功读取数据，则执行该case处理语句
		fmt.Printf("2th case is selected. e2=%v",e2)
	case <- time.After(2 * time.Second):
		fmt.Println("Timed out")
	}

}

func main011() {
	var t *time.Timer

	f := func() {
		fmt.Printf("Expiration time : %v.\n", time.Now())
		fmt.Printf("C`s len: %d\n", len(t.C))
	}

	t = time.AfterFunc(1 * time.Second, f)

	time.Sleep(2 * time.Second)
}

func main012()  {
	var ticker *time.Ticker = time.NewTicker(1 * time.Second)

	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Second * 5)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

func main() {
	var ticker *time.Ticker = time.NewTicker(1 * time.Second)

	// number 为指定的执行次数
	num := 2
	c := make(chan int, num)

	go func() {
		fmt.Println("start go func")
		for t := range ticker.C {
			c <- 1
			fmt.Println("Tick at", t)
		}
	}()

	//for cc := range c {
	//	fmt.Println(cc)
	//}
	time.Sleep(time.Second * 5)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}