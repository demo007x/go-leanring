package main

import "fmt"

type Property struct {
	value int
}

// 设置属性的值
func (p *Property) SetValue(v int)  {
	p.value = v
}

func (p *Property) getValue() int  {
	return p.value
}
func main() {
	//p := &Property{}
	p := new(Property)
	p.SetValue(1001)
	fmt.Println(p.getValue())
}

