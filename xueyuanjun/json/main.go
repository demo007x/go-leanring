package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
	Website string
	Age int
	Male bool
	Skills []string
}

func main() {
	user := User{
		"学院君",
		"https://xueyuanjun.com",
		18,
		true,
		[]string{"Golang", "PHP", "C", "Java", "Python"},
	}

	// 编码
	u, err := json.Marshal(user)
	if err != nil {
		fmt.Println("json 编码失败： %v", err)
		return
	}

	fmt.Printf("json 格式数据： %s\n", u)
	
	var user2 User
	
	err = json.Unmarshal(u, &user2)

	if err != nil {
		fmt.Printf("JSON 解码失败：%v\n", err)
		return
	}

	fmt.Printf("JSON 解码结果: %#v\n", user2)
}
