package main

import "fmt"

func hello2(ch chan int)  {
	ch<-3
}

func main()  {
	ch := make(chan int)

	go hello2(ch);
	fmt.Println(<-ch)
}
