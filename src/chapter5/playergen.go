package main

import (
	"bytes"
	"fmt"
)

//func playerGen(name string) func() (string, int)  {
//	// 血量
//	hp := 150
//
//	return func() (string, int) {
//		return  name, hp
//	}
//}
//
//func main() {
//	generator := playerGen("high noon")
//	name , hp := generator()
//
//	fmt.Println(name, hp)
//}

func joinString(slist ...string) string {
	var b bytes.Buffer
	for _, s := range slist {
		b.WriteString(s)
	}

	return b.String()
}

func main() {
	fmt.Println(joinString("pig ", "and", "rat"))
	fmt.Println(joinString("hammer", " mom", " and", " hawk"))
}
