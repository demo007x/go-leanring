package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main01() {
	path, err := os.Getwd()
	filePath := path + "/os/main.go"
	if err != nil {
		panic(err)
	}
	fmt.Println(path)
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println(string(fileContent))
}

func main002() {
	path, _ := os.Getwd()
	filePath := path + "/os/main.go"
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(file))
}

func main003() {
	path, _ := os.Getwd()
	dir, err := ioutil.ReadDir(path)
	pathInfo := map[string]interface{}{}
	if err != nil {
		panic(err)
	}
	dirMap := map[int]interface{}{}
	for index, file := range dir {
		pathInfo["fileName"] = file.Name()
		pathInfo["isDir"] = file.IsDir()
		pathInfo["sys"] = file.Sys()
		pathInfo["modTime"] = file.ModTime().Format("2006-01-02 15:04:05")
		pathInfo["mode"] = file.Mode()
		//pathInfo["content"], _ = ioutil.ReadAll(file)
		dirMap[index] = pathInfo
	}

	json, err := json.MarshalIndent(dirMap, "", "  ")
	if err != nil {
		panic(err)
	}

	//fmt.Println(string(json))
	// 输出json文件
	path, _ = os.Getwd()
	jsonFile := path + "/os/dir.json"
	//jsonFileHand, err := os.Create(jsonFile)
	jsonFileHand, err := os.Create(jsonFile)
	if err != nil {
		panic(err)
	}
	defer jsonFileHand.Close()
	n, err := jsonFileHand.Write([]byte(json))
	fmt.Println(n)
}

func main004() {
	path, _ := os.Getwd()
	jsonFile := path + "/os/dir.json"
	fileHandle, err := os.Open(jsonFile)
	if err != nil {
		panic(err)
	}
	defer fileHandle.Close()
	//fmt.Println(fileHandle)
	content, err := ioutil.ReadAll(fileHandle)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}

func main() {
	//pwd, _ := os.Getwd()
	//jsonFile := pwd + "/os/dir.json"

	//fmt.Println(path.Dir(jsonFile))
	//fmt.Println(path.Base(jsonFile))
	//fmt.Println(path.Ext("dir.json"))
	//fmt.Println(path.IsAbs(jsonFile))
	//fmt.Println(path.IsAbs("./os/dir.json"))
	//fmt.Println(path.Split("./os/dir.json"))
	//fmt.Println(path.Join("./os", "dir.json"))
	//fmt.Println(path.Clean("../../go-leanring/./os/dir.json"))

	fmt.Println(filepath.Walk("./os", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(info.Name(), info.IsDir())
		return nil
	}))
}
