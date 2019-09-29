package main

type Cat struct {
	Color string
	Name string
}

type BlackCat struct {
	Cat
}

// 构造基类
func NewCat(name string) *Cat {
	return &Cat{
		Name:  name,
	}
}

// 构造子类
func NewBlackCat(color string) *BlackCat  {
	cat := &BlackCat{}
	cat.Color = color
	return cat
}
