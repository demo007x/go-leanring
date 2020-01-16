package main

import (
	"fmt"
)

type Animaler interface {
	// 获取名字
	Get() string
	// 定义吃的方法
	Eat(food string)
	// 定义走的方法
	Walk()
}

type User struct {
	uName string
	uAge  int
	uSex  string
}

// User 实现Animaler的接口
func (u User) Get() string {
	return u.uName
}

func (u User) Eat(food string) {
	fmt.Println(food)
}

func (u User) Walk() {
	fmt.Println("User is walk")
}

type Bird struct {
	Name     string
	Color    string
	Category string
}

func (b Bird) Get() string {
	return b.Name
}

func (b Bird) Eat(food string) {
	fmt.Println("this bird is eat ", food)
}

func (b Bird) Walk() {
	fmt.Println("the bird is walking")
}

// 获取名字的方法
func getName(a Animaler) string {
	return a.Get()
}

func main() {
	user := User{
		uName: "test",
		uAge:  1,
		uSex:  "man",
	}

	bird := &Bird{
		Name:     "duck",
		Color:    "blue",
		Category: "bird",
	}
	fmt.Println(user)
	fmt.Println(bird)
	fmt.Println("=================")
	fmt.Println(getName(user))
	fmt.Println(getName(bird))
	fmt.Println("=================")

	print(user)

	uu := new(User)
	uu.uName = "demo"
	print(uu)
}

func print(resource interface{}) {
	switch v := resource.(type) {
	case Animaler:
		fmt.Println("Animaler", v.Get())
	default:
		fmt.Println("default")
	}
}
