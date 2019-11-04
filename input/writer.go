package main

import (
	"bytes"
	"fmt"
)

//func main() {
//	reader := strings.NewReader("Go 语音中文网")
//	p := make([]byte, 6)
//	n, err := reader.ReadAt(p, 2)
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Printf("%s, %d\n", p, n)
//}

func main() {
	//file, err := os.Create("writeAt.txt")
	//if err != nil {
	//	panic(err)
	//}
	//
	//defer file.Close()
	//file.WriteString("Golang中文社区-这里是多余的")
	//n, err := file.WriteAt([]byte("Go语言中文网"), 24)
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(n)
	//file, err := os.Open("writeAt.txt")
	//if err != nil {
	//	panic(err)
	//}
	//
	//defer file.Close()
	//writer := bufio.NewWriter(os.Stdout)
	//writer.ReadFrom(file)
	//writer.Flush()

	//reader := bytes.NewReader([]byte("go语音中文版"))
	//reader.WriteTo(os.Stdout)

	//reader := strings.NewReader("Go语言中文版")
	//reader.Seek(-6, io.SeekEnd)
	//r, _, _ := reader.ReadRune()
	//fmt.Printf("%c\n", r)

	var ch byte
	fmt.Scanf("%c\n", &ch)
	buffer := new(bytes.Buffer)
	err := buffer.WriteByte(ch)
	if err == nil {
		fmt.Println("写入一个字节成功！准备读该字节。。。。。。")
		newch, _ := buffer.ReadByte()
		fmt.Printf("读取的字节：%c\n", newch)
	} else {
		fmt.Println("写入错误")
	}
}



