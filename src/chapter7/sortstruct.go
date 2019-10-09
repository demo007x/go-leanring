package main

import (
	"fmt"
	"sort"
)

type HeroKind int

const (
	None HeroKind = iota
	Tank
	Assassin
	Mage
)

type Hero struct {
	Name string
	Kind HeroKind
}

type Heros []*Hero

// 实现sort.Interface 接口比较元素方法
//func (s Heros) Len() int  {
//	return len(s)
//}
//
//func (s Heros) Less(i, j int) bool {
//	if s[i].Kind != s[j].Kind {
//		return s[i].Kind < s[j].Kind
//	}
//
//	return s[i].Name < s[j].Name
//}
//
//func (s Heros) Swap(i, j int)  {
//	s[i], s[j] = s[j], s[i]
//}

func main() {
	// 准备英雄列表
	heros := Heros{
		&Hero{
			Name: "吕布",
			Kind: Tank,
		},
		&Hero{
			Name: "李白",
			Kind: Assassin,
		},
		&Hero{
			Name:"妲己",
			Kind: Mage,
		},
		&Hero{
			Name: "貂蝉",
			Kind: Assassin,
		},
		&Hero{
			Name: "关羽",
			Kind: Tank,
		},
		&Hero{
			Name: "诸葛亮",
			Kind: Mage,
		},
	}

	sort.Slice(heros, func(i, j int) bool {
		if heros[i].Kind != heros[j].Kind {
			return heros[i].Kind < heros[j].Kind
		}

		return heros[i].Name < heros[j].Name
	})

	for _, v := range heros {
		fmt.Printf("%+v\n", v)
	}
}

