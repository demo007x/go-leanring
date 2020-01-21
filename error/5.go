package main

import "fmt"

// 自定义错误接口
type MyError struct {
	Msg string
	Err error
}

func (me *MyError) Error() string {
	return me.Msg
}

type Geter interface {
	Get() string
}

type User struct {
	Name string
	Age  int
}

func (u User) Get() string {
	return u.Name
}

func main() {
	user := &User{
		Name: "test",
		Age:  20,
	}

	name, err := getName(user)
	if err != nil {
		switch err.(type) {
		case *MyError:
			fmt.Println("MyError")
			fmt.Println(err)
		default:
			fmt.Println("not found the error")
		}
	} else {
		fmt.Println(name)
	}
}

func getName(u Geter) (name string, err error) {
	if u.Get() == "demo" {
		return u.Get(), nil
	}

	return "", &MyError{
		Msg: "not found the value",
	}
}
