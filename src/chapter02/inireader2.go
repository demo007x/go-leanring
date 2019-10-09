package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getIniFileValues(filepath string) map[string]map[string]string {
	values := make(map[string]map[string]string)

	file, err := os.Open(filepath)

	if nil != err {
		return values
	}
	defer file.Close()
	fileReader := bufio.NewReader(file)

	for {
		// 按换行获取字符串
		str, err := fileReader.ReadString('\n')
		// 如果读取错误， 则继续下一行
		if nil != err {
			break
		}
		// 如果是注释行， 则跳过
		if str[0] == ';' {
			continue
		}

		// 如果为空行， 则跳过
		str = strings.TrimSpace(str)
		if str == "" {
			continue
		}

		// 解析块
		var sectionName string
		if str[0] == '[' && str[len(str)-1] == ']' {
			sectionName = str[1 : len(str)-1]
		}

		// 解析等号
		pair := strings.Split(str, "=")
		if len(pair) == 2 {
			// 获取键值对
			key := strings.TrimSpace(pair[0])
			value := strings.TrimSpace(pair[1])
			if values[sectionName] == nil {
				values[sectionName] = map[string]string{}
			}
			values[sectionName][key] = value
		}
	}

	return values
}

func main() {
	values := getIniFileValues("config.ini")

	//fmt.Printf("ini file's value is %#v", values)

	for sectionName, val := range values {
		for k, v := range val {
			fmt.Println(sectionName, k, v)
		}
	}
}
