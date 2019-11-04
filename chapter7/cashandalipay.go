package main

import "fmt"

type Alipay struct {

}

func (a *Alipay) CanUseFaceID()  {

}

type Cash struct {

}
func (a *Cash) Stolen() {

}
// 具备刷脸特性的借口
type CantainCanUserFaceID interface {
	CanUseFaceID()
}

type ContainStolen interface {
	Stolen()
}

// 打印支付方式具备的特点
func print(payMethod interface{})  {
	switch payMethod.(type) {
	case CantainCanUserFaceID:
		fmt.Println("%T can use faceid\n", payMethod)
	case ContainStolen:
		fmt.Printf("%T may be stolen\n", payMethod)
	}
}

func main() {
	print(new(Alipay))
	print(new(Cash))
}