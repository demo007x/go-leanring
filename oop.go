package main

import (
	"fmt"
	"math"
)

type IAnimal interface {
	GetName() string
	Call() string
	FavorFood() string
}

func main() {
	fmt.Println(math.MaxInt32)
	//var animal = animal2.Animal{"小狗"}
	//var ianimal IAnimal = animal2.Dog{&animal}
	//if dog, ok := ianimal.(animal2.Animal); ok {
	//	fmt.Println(dog.GetName())
	//} else {
	//	fmt.Println(false)
	//}
}
