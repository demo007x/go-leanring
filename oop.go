package main

import (
	"fmt"
	animal2 "go-leanring/animal"
)

func main() {
	ani := animal2.NewAnimal("狗狗")
	dog := animal2.Dog{ani}
	fmt.Println(dog.GetName(), dog.Call(), dog.FavorFood())
}
