package main

type Point struct {
	X int
	Y int
}
type Player struct {
	Name string
	HealthPoint int
	MagicPoint int
}

type Command struct {
	Name string
	Var *int // 指令绑定的变量
	Comment string
}

type People struct {
	name string
	child *People
}

var p Point // 实例化
var version int = 1

func main() {
	p.X = 10
	p.Y = 20

	tank := new(Player)
	tank.Name = "Canon"
	tank.HealthPoint = 300
	tank.MagicPoint = 500
	//++++++//
	cmd := &Command{}
	cmd.Name = "version"
	cmd.Var = &version
	cmd.Comment = "show version"
	//++++++//
	
	relation := &People{
		name:  "爷爷",
		child: &People{
			name:  "爸爸",
			child: &People{
				name:  "Me",
				child: nil,
			},
		},
	}
}
