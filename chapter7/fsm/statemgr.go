package main

import "github.com/anziguoer/go-leanring/pkg/mod/github.com/pkg/errors@v0.8.1"

type StateManager struct {
	// 已近添加的状态
	stateByName map[string]State
	// 状态改变时的回调
	OnChange func(from, to State)
	// 当前状态
	curr State
}
// 根据名字获取指定的状态
func (sm *StateManager) Get(name string) State {
	if v, ok := sm.stateByName[name]; ok {
		return v
	}

	return nil
}

// 添加一个状态到管理器中
func (sm *StateManager) Add(s State)  {
	// 获取状态的名称
	name := s.Name()
	// 将s转换为能设置名字的接口，然后调用该借口
	s.(interface{
		setName(name string)
	}).setName(name)
	// 根据状态名获取已近添加的状态， 检测该状态是否存在
	if sm.Get(name) != nil {
		panic("duplicate state:"+name)
	}
	// 根据名字保存到map中
	sm.stateByName[name] = s
}

// 初始化状态管理器
func NewStateManager() *StateManager {
	return &StateManager{
		stateByName: make(map[string]State),
	}
}

// 状态没有找到的错误
var ErrStateNotFound  = errors.New("state not found")
// 禁止在同状态见转移的错误
var ErrCannotTransitToState = errors.New("cannot transit to state")
// 获取当前状态
func (sm *StateManager) CurrState() State  {
	return sm.curr
}

// 当前状态能否转移到目标状态
func (sm *StateManager) CanCurrTransitTo(name string) bool {
	if sm.curr == nil {
		return true
	}
	// 相同的状态不能转换
	if sm.curr.Name() == name && !sm.curr.EnableSameTransit() {
		return false
	}

	return sm.curr.CanTransitTo(name)
}

func (sm *StateManager) Transit(name string) error {
	next := sm.Get(name)
	if next == nil {
		return ErrStateNotFound
	}
	// 记录转移前的状态
	pre := sm.curr
	if sm.curr != nil {
		// 相同状态不用转移
		if sm.curr.Name() == name && !sm.curr.EnableSameTransit() {
			return ErrCannotTransitToState
		}
		// 不能转移到目标状态
		if !sm.curr.CanTransitTo(name) {
			return ErrCannotTransitToState
		}
		sm.curr.OnEnd()
	}
	// 将当前状态切换为要转移的目标状态
	sm.curr = next
	// 调用新状态的开始
	sm.curr.OnBegin()
	if sm.OnChange != nil {
		sm.OnChange(pre, sm.curr)
	}
	return nil
}
