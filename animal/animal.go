package animal

type Animal struct {
	name string
}

func (a Animal) Call() string {
	return "动物的叫声..."
}

func (a Animal) FavorFood() string {
	return "动物爱吃的食物..."
}

func (a Animal) GetName() string {
	return a.name
}

func NewAnimal(name string) *Animal {
	return &Animal{name: name}
}
