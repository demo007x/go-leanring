package main

import (
	"fmt"
	"reflect"
)

type Brand struct {
}

func (t Brand) show() {

}

// 别名
type FakeBrand = Brand

type Vehicle struct {
	FakeBrand
	Brand
}

func main() {
	var a Vehicle
	a.FakeBrand.show()
	ta := reflect.TypeOf(a) // 获取a的放射对象

	for i := 0; i < ta.NumField(); i++ {
		f := ta.Field(i)
		fmt.Printf("FiledName: %v, FieldType: %v\n", f.Name, f.Type.Name())
	}
}
