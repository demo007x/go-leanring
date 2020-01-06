package main

import "fmt"

//在 Go 语言中，并非在调用延迟函数的时候才确定实参，而是当执行 defer 语句的时候，就会对延迟函数的实参进行求值
func printA(a int)  {
	fmt.Println("value oof a in deferred function", a)
}

//从上面的输出，我们可以看出，在调用了 defer 语句后，虽然我们将 a 修改为 10，但调用延迟函数 printA(a)后，仍然打印的是 5。
func main() {
	a := 5
	defer printA(a)
	a = 10

	fmt.Println("value of a before deferred function call", a)
}
