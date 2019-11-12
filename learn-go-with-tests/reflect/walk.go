package main

import (
	"fmt"
	"reflect"
)

func Walk(x interface{}, fn func(input string)) {
	// 使用放射
	//val := reflect.ValueOf(x) //返回一个给定变量的value
	//// 指针类型的不能使用 Numfield 来提取
	//if val.Kind() == reflect.Ptr {
	//	val = val.Elem() // 使用elem 来提取指针底层的值
	//}

	val := getValue(x)
	//numberOfValues := 0
	//var getField func(int) reflect.Value

	walkValue := func(value reflect.Value) {
		Walk(value.Interface(), fn)
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			//numberOfValues = val.NumField()
			//getField = val.Field
			//Walk(field.Interface(),fn)
			for i := 0; i < val.NumField(); i++ {
				walkValue(val.Field(i))
			}
		case reflect.Slice, reflect.Array:
			// 如果是切片， 需要遍历切片的每一个值
			//for i :=0 ; i < val.NumField() ; i++ {
			//	Walk(val.Index(i).Interface(), fn)
			//}
			//numberOfValues = val.Len()
			//getField = val.Field

			for i := 0; i < val.NumField(); i++ {
				walkValue(val.Index(i))
			}

		case reflect.Map:
			for _, key := range val.MapKeys(){
				//Walk(val.MapIndex(key), fn)
				walkValue(val.MapIndex(key))
			}
		}

		//for i := 0; i < numberOfValues; i++ {
		//	Walk(getField(i).Interface(), fn)
		//}
	}
}

func getValue(x interface{}) reflect.Value  {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}

func main() {
	x := struct {
		name string
	}{
		"hello",
	}
	val := reflect.ValueOf(x)
	field := val.Field(0)
	fmt.Printf("val value is %#v \n, field value is %#v", val, field)
}
