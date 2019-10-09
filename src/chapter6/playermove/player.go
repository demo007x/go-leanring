package main

type Player struct {
	currPos   Vec2
	targetPos Vec2
	speed     float32
}

// 玩家移动的目标位置
func (p *Player) MoveTo(v Vec2) {
	p.targetPos = v
}

// 获取当前的位置
func (p *Player) Pos() Vec2 {
	return p.currPos
}

// 判断是否到达了目标位置
func (p *Player) IsArrived() bool {
	return p.currPos.DistanceTo(p.targetPos) < p.speed
}

func (p *Player) Update() {
	if !p.IsArrived() {
		dir := p.targetPos.Sub(p.currPos).Normalize()
		newPos := p.currPos.Add(dir.Scale(p.speed))
		p.currPos = newPos
	}
}

func newPlayer(speed float32) *Player {
	return &Player{
		speed: speed,
	}
}
