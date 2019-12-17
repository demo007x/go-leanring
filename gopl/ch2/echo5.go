package main

import "fmt"

func main() {
	//a := int(5)
	//b := int(4)
	//
	//fmt.Println(a / b)
	//var c, d float32
	//c = 5
	//d = 4
	//
	//fmt.Println(c / d)

	//一个算术运算的结果，不管是有符号或者是无符号的，如果需要更多的bit位才能正确表示的话，就说明计算结果是溢出了。
	//超出的高位的bit位部分将被丢弃。如果原始的数值是有符号类型，而且最左边的bit位是1的话，那么最终结果可能是负的，
	//例如int8的例子：
	//var u uint8 = 255
	//fmt.Println(u, u + 1, u + u) // 255 0 1
	//
	//var i int8 = 127
	//fmt.Println(i, i + i, i*i)

	// 位操作

	var a uint8 = 1

	fmt.Printf("a %08b \n", a)
	fmt.Printf("1<<1 => %08b -> %d \n", 1<<1, 1<<1)
	fmt.Printf("1<<5 => %08b -> %d \n", 1<<5, 1<<5)
	fmt.Printf("1<<2 => %08b -> %d \n", 1<<2, 1<<2)
	fmt.Printf("2<<2 => %08b -> %d \n", 2<<2, 2<<2)
	fmt.Printf("1>>2 => %08b -> %d \n", 1>>2, 1>>2)

	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b \n", x)
	fmt.Printf("%08b \n", y)

	fmt.Printf("%08b \n", x&y)
	fmt.Printf("%08b \n", x|y)
	fmt.Printf("%08b \n", x^y)
	fmt.Printf("%08b \n", x&^y)

	for i := uint(0);  i< 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println(i) //1, 5
		}
	}

	fmt.Printf("%08b \n", x<<1)
	fmt.Printf("%08b \n", x>>1)

	// 左移操作 << 值的计算： 值 * 2^n, 例如：2<<2 值为： 2*2^2 = 8
	// 右移操作 >> 值的计算： 值/ 2^n， 例如：2 >> 2， 值为： 2 / 2^2 =

	medals := []string{"gold", "silver", "bronze"}
	for i := len(medals) -1 ; i >= 0; i-- {
		fmt.Println(medals[i])
	}

	var apples int32 = 1
	var oranges  int16 = 2
	var compote  = apples + oranges // 编译不通过
	fmt.Println(compote)
}
