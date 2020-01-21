package main

import (
	"fmt"
	animal2 "go-leanring/animal"
)

type IAnimal interface {
	GetName() string
	Call() string
	FavorFood() string
}

func main() {
	var animal = animal2.Animal{"小狗"}
	var ianimal IAnimal = animal2.Dog{&animal}
	if dog, ok := ianimal.(animal2.Animal); ok {
		fmt.Println(dog.GetName())
	} else {
		fmt.Println(false)
	}
}
