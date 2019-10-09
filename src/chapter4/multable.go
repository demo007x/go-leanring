package main

import "fmt"

func main() {
	for y := 1; y <= 9; y++ {
		for x := 1; x <= y; x++ {
			if x == 1 {
				fmt.Printf("%d*%d = %d", y, x, x*y)
			} else {
				fmt.Printf("\t %d*%d = %d", y, x, x*y)
			}
		}
		fmt.Println()
	}

	for key, value := range []int{1, 2, 3, 4} {
		fmt.Printf("key:%d value:%d \n", key, value)
	}

	str := "hello 你好"
	for key, value := range str {
		fmt.Printf("key:%d value:0x%x \n", key, value)
	}

	m := map[string]int{
		"hello": 100,
		"world": 200,
	}

	for key, value := range m {
		fmt.Println(key, value)
	}

	c := make(chan int)

	go func() {
		c <- 1
		c <- 2
		c <- 3
		c <- 4
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}
