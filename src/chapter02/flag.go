package main

import (
	"fmt"
	"unicode/utf8"
)

//var mode  = flag.String("mode","--help", "process mode")
func main() {
	//flag.Parse()
	//fmt.Printf("%T \n", mode)
	//fmt.Println(*mode)
	tip1 := "genji is a ninja"
	fmt.Println(len(tip1))

	tip2 := "杨玉龙"
	fmt.Println(len(tip2))

	fmt.Println(utf8.RuneCountInString(tip2))
	fmt.Println(utf8.RuneCountInString("龙忍出鞘, fight!"))

	theme := "狙击 start"

	for i := 0; i < len(theme); i++  {
		fmt.Printf("ascill: %c %d\n", theme[i], theme[i])
	}

	for _, value := range theme {
		fmt.Printf("Unicode: %c %d\n", value, value)
	}
}
