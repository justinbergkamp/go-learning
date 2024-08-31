package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name, language string) string {

	if name == "" {
		name = "World"
	}
	r := greetingPrefix(language) + name
	return r
}

func greetingPrefix(language string) string {
	helloPrefix := englishHelloPrefix

	switch language {
	case spanish:
		helloPrefix = spanishHelloPrefix
	case french:
		helloPrefix = frenchHelloPrefix
	}
	return helloPrefix
}

func main() {
	fmt.Println(Hello("World", "English"))
}
