package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getValue(filename, expectSection, expectKey string) string {
	file, err := os.Open(filename)
	if err != nil {
		return ""
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var sectionName string
	var line int
	for {
		line++
		linestr, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		linestr = strings.TrimSpace(linestr)
		// 忽略行左右的空行
		if linestr == "" {
			continue
		}

		// 忽略注释
		if linestr[0] == ';' {
			continue
		}
		// 读取键和键值的代码
		if linestr[0] == '[' && linestr[len(linestr)-1] == ']' {
			sectionName = linestr[1 : len(linestr)-1]
			fmt.Println(sectionName, line)
		}

		if sectionName == expectSection {
			pair := strings.Split(linestr, "=")
			if len(pair) == 2 {
				key := strings.TrimSpace(pair[0])
				if key == expectKey {
					return strings.TrimSpace(pair[1])
				}
			}
		}
	}
	return ""
}

func main() {
	params := getValue("./config.ini", "remote \"origin\"", "fetch")
	fmt.Println(params)

	fmt.Println(getValue("./config.ini", "core", "hideDotFiles"))
}
