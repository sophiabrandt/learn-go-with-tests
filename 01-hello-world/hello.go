package main

import "fmt"

const helloPrefixEN = "Hello, "

func Hello(name string) string {
	return helloPrefixEN + name
}

func main() {
	fmt.Println(Hello("world"))
}
