package main

import "fmt"

func main() {
	var ch chan int
	ch = make(chan int, 5)
	for i := 0; i < 5; i++ {
		ch <- i
	}

	close(ch)

	//for {
	//	var b int
	//	b, ok := <-ch
	//	if ok == false {
	//		fmt.Println("chan is close")
	//		break
	//	}
	//	fmt.Println(b)
	//}

	for i := 0; i < 6; i++  {
		fmt.Println(<-ch)
	}

}
