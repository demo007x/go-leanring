package main

import "fmt"

type Dictionary struct {
	data map[interface{}]interface{}
}
// 根据建获取值
func (d *Dictionary) Get(key interface{}) interface{} {
	return d.data[key]
}
// 设置健值
func (d *Dictionary) Set(key, val interface{})  {
	d.data[key] = val
}

// 遍历所有的健值，如果返回的是false，停止遍历
func (d *Dictionary) Visit(callback func(k, v interface{}) bool) {
	if callback == nil {
		return
	}

	for k, v := range d.data {
		if !callback(k, v) {
			return
		}
	}
}
// 清空所有的数据
func (d *Dictionary) Clear() {
	d.data = make(map[interface{}]interface{})
}

// 创建一个字典
func NewDictionary() *Dictionary {
	d := &Dictionary{}
	d.Clear()

	return d
}

func main() {
	dict := &Dictionary{}

	dict.Set("My Factory", 60)
	dict.Set("Terra", 36)
	dict.Set("Dont't Hungry", 24)
	// 获取值及打印值
	favorite := dict.Get("Terra Craft")
	fmt.Println(favorite)

	// 遍历所有的字典元素
	// dict.Visit(func(k, v interface{}) bool {
	// 	if v.(int) > 40 {
	// 		// 输出很贵
	// 		fmt.Println(k, "is expensive")
	// 		return true
	// 	}
	//
	// 	fmt.Println(k, "is cheap")
	// 	return true
	// })
}

