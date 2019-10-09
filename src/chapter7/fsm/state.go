package main

import "reflect"

type State interface {
	// 获取状态的名字
	Name() string
	// 该状态是否允许在相同状态之间转换
	EnableSameTransit() bool
	// 响应状态开始时
	OnBegin()
	// 响应状态结束时
	OnEnd()
	// 判断能否转移到其他的状态
	CanTransitTo(name string) bool
}

func StateName(s State) string {
	if s == nil {
		return "none"
	}
	// 使用放射获取状态的名称
	return reflect.TypeOf(s).Elem().Name()
}
