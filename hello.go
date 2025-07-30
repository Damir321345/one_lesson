package main

import "fmt"

const enHelloPrefix = "Hello "
const spanish = "Spanish"
const spanishHelloPrefix = "Hola "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	if language == spanish {
		return spanishHelloPrefix + name
	}
	return enHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Freed", "Spanis"))
}
