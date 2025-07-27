package main

import "fmt"

const englHelloPregix = "Hello "

func Hello(name string) string {
	return englHelloPregix + name
}

func main() {
	fmt.Println(Hello("Любое значение"))
}
