package eventsys

import "fmt"

type Actor struct {
	
}

func (a *Actor) OnEvent(param interface{})  {
	fmt.Println("actor event:", param)
}

func GlobalEvent(param interface{})  {
	fmt.Println("global event:", param)
}

func main() {
	a := new(Actor)
	RegisterEvent("OneSkill", a.OnEvent)
	RegisterEvent("OnSkill", GlobalEvent)

	CallEvent("OnSkill", 100)
}


