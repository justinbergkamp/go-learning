package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	r := englishHelloPrefix + name
	return r
}

func main() {
	fmt.Println(Hello("World"))
}
