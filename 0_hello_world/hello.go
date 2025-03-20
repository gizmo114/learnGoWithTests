package main

import "fmt"

const spanish = "Spanish"

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
)

func main() {
	fmt.Println(Hello("world", ""))
}

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	if language == spanish {
		return spanishHelloPrefix + name + "!"
	}
	return englishHelloPrefix + name + "!"
}
