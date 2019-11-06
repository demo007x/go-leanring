package main

import "fmt"

const spanish = "Spanish"
const helloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name, language string) string {
	// if name == "" {
	// 	name = "world"
	// }

	// if language == "Spanish" {
	// 	return spanishHelloPrefix + name
	// }

	prefix := helloPrefix

	switch language {
	case spanish:
		prefix = spanish
	case spanishHelloPrefix:
		prefix = spanishHelloPrefix
	default:
		prefix = helloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
