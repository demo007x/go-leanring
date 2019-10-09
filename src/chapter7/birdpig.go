package main

import "fmt"

type Flyer interface {
	Fly()
}

type Walker interface {
	Walk()
}

// 定义鸟类
type Bird struct {

}

func (b *Bird) Fly()  {
	fmt.Println("bird: fly")
}

type Pig struct {

}

func (p *Pig) Walk() {
	fmt.Println("pig: walk")
}

func main() {
	animals := map[string]interface{}{
		"bird": new(Bird),
		"pig": new(Pig),
	}
	// 遍历映射
	for name, obj := range animals {
		// 判断对象是否是飞行动物
		f,isFlyer := obj.(Flyer)
		// 判断对象是否是行走动物
		w, isWalker := obj.(Walker)

		fmt.Printf("name:%s isFlayer: %v isWalker:%v\n", name,isFlyer, isWalker)

		if isFlyer {
			f.Fly()
		}

		if isWalker {
			w.Walk()
		}
	}
}
