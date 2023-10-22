package main

import "fmt"

const helloPrefixEN = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return helloPrefixEN + name
}

func main() {
	fmt.Println(Hello("world"))
}
