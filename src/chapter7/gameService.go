package main

type Service interface {
	Start()
	Log(string)
}

type Logger struct {
	
}

func (g *Logger) Log(l string) {
	
}

type GameService struct {
	Logger
}

func (g *GameService) Start()  {
	
}
