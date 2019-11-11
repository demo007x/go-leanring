package main

type foo struct {
	bar int
}

func main() {
	// 第一题： 编译不能通过
	// var f foo
	// f.bar, tmp := 1,2
	//解答：：= 操作符不能给结构体赋值
	// 第二题：
	fmt.Println(~2) //invalid character U+007E '~'
	// 很多语言都是采用 ~ 作为按位取反运算符，Go 里面采用的是 ^ 。按位取反之后返回一个每个 bit 位都取反的数，对于有符号的整数来说，是按照补码进行取反操作的（快速计算方法：对数 a 取反，结果为 -(a+1) ），对于无符号整数来说就是按位取反
}
