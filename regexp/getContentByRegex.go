package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	resp, err := http.Get("https://baidu.com")
	if err != nil {
		fmt.Println("http get error.")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	src := string(body)
	// 将html标签转换为小写
	re, err := regexp.Compile("\\<[\\S\\s]+?\\>")
	if err != nil {
		fmt.Println(err)
		return
	}
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	// 去除样式
	re, err = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	if err != nil {
		fmt.Println(err)
	}
	src = re.ReplaceAllString(src, "")

	// 去除script标签
	re, err = regexp.Compile("\\<script[\\S\\s]+?\\</script>")
	if err != nil {
		fmt.Println(err)
		return
	}
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")

	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")

	fmt.Println(strings.TrimSpace(src))
}
