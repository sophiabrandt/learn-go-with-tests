package main

import "fmt"

const (
	languageES    = "es"
	languageFR    = "fr"
	helloPrefixEN = "Hello, "
	helloPrefixES = "Hola, "
	helloPrefixFR = "Bonjour, "
)

func greetingPrefix(language string) string {
	var prefix string
	switch language {
	case "en":
		prefix = helloPrefixEN
	case "es":
		prefix = helloPrefixES
	case "fr":
		prefix = helloPrefixFR
	default:
		prefix = helloPrefixEN
	}
	return prefix
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := greetingPrefix(language)
	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
