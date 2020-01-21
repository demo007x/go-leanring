package main

import "fmt"

type tempInterface interface {
	tmpFunc()
	tmpFuncPtr()
}

type temStruct struct {
}

func (ts temStruct) tmpFunc() {
	fmt.Println("functmpFunc")
}

func (ts *temStruct) tmpFuncPtr() {
	fmt.Println("functemFuncPtr")
}

func main() {
	// 这里如果用了new,则得到的是tmpStruct的指针,
	var tt tempInterface = new(temStruct)
	tt.tmpFunc()
	tt.tmpFuncPtr()

	var tt2 tempInterface = temStruct{}
	tt2.tmpFunc()
	tt2.tmpFuncPtr()
}
