package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func mustCompile(str string) *regexp.Regexp {
	regexp, err := regexp.Compile(str)
	if err != nil {
		panic(`regexp: Compile(` + strconv.Quote(str) + `):` + err.Error())
	}

	return regexp
}

func main() {
	defer fmt.Println("宕机后要做的事情1")
	defer fmt.Println("宕机后要做的事情2")
	panic("宕机")
}
