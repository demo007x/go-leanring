package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	dir := os.Args[1] // 获取命令行下面的目录参数
	listAll(dir, 0)
}

func listAll(dir string, curHier int)  {
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, info := range fileInfos {
		// 如果是目录的情况
		if info.IsDir() {
			for tmpHier := curHier; tmpHier > 0; tmpHier--  {
				fmt.Printf("|\t")
			}
			fmt.Println(info.Name(), "\\")
			listAll(dir + "\\"+info.Name(), curHier + 1)
		} else {
			for tmpHier := curHier; tmpHier > 0; tmpHier-- {
				fmt.Println("|\t")
			}
			fmt.Println(info.Name())
		}
	}
}
