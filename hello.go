package main

import "fmt"

func main() {
	fmt.Printf(Hello("world"))
}

const helloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return helloPrefix + name
}
