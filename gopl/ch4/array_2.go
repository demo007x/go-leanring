package main

import "fmt"

func main() {
	s := "Hello, world!"
	ss := s[1:3]

	fmt.Println(s, ss)

	var runes []rune
	for _,r := range "Hello, 世界" {
		runes = append(runes, r)
	}

	fmt.Printf("%q\n", runes)
}
