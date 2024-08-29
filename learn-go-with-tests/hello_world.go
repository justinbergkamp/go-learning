package main

import "fmt"

func Hello(name string) string {
	r := "Hello, " + name
	return r
}

func main() {
	fmt.Println(Hello("World"))
}
