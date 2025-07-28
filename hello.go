package main

import "fmt"

const enHelloPrefix = "Hello "

func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return enHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Rot"))
}
