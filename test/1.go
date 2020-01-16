package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) String() string {
	return u.Name
}

var u User

// 引用类型的demo
func main() {
	u = User{
		Name: "a",
		Age:  1,
	}
	fmt.Println(u)
}
