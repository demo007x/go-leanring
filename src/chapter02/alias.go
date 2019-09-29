package main

import (
	"fmt"
	"time"
)

type NewInt int

type IntAlias = int

// 别名
type MyDuration time.Duration

func (m MyDuration) EasySet(a string)  {
	
}

func main() {
	var a NewInt

	fmt.Printf("a type: %T\n", a)

	var a2 IntAlias

	fmt.Printf("a2 type: %T\n", a2)
}
