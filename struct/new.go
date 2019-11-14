package main

import "fmt"

type Persion struct {
	Name string
	Age int
}
// go 语言中一般通过工厂模式实例化结构体
func NewPersion(name string, age int) *Persion {
	return &Persion {
		name,
		age,
	}
}
func main (){
	var p1 Persion
	p1.Name = "test"
	p1.Age = 21
	fmt.Printf("p1 %#v \n", p1)
	p1.Name = "test_p1"
	fmt.Println("p1 name = " + p1.Name)
	p2 := &Persion{
		"anziguoer",
		18,
	}
	fmt.Printf("p2 %#v \n", p2)
	fmt.Println("p2.name = " + p2.Name)
	p3 := Persion{
		"xxx",
		20,
	}
	fmt.Printf("p3 %#v \n", p3)
	fmt.Println("p3 name = " + p3.Name)
	p3.Name = "haha"
	fmt.Println("p3 name = " + p3.Name)

	np := NewPersion("longer", 30)
	fmt.Printf("np value is %#v \n", np)
	np.Name = "anzi" + np.Name 
	fmt.Printf("np value is %#v \n", np)
}