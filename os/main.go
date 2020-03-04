package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	pwd, e := os.Getwd()
	if e != nil {
		log.Fatal(e)
	}
	file, err := os.Open(pwd + "/os/main.go")
	if err != nil {
		log.Fatal(err)
	}
	bytes, _ := ioutil.ReadAll(file)
	fmt.Println(string([]byte(bytes)))
}
