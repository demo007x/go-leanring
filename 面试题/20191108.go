package main

type foo struct {
	bar int
}

func main() {
	// 第一题： 编译不能通过
	// var f foo
	// f.bar, tmp := 1,2

	// 第二题：
	fmt.Println(~2) //invalid character U+007E '~'
}
