package main

import "fmt"

func main() {
	fmt.Printf(Hello("world", ""))
}

const french = "French"
const spanish = "Spanish"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Elodie, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {

	switch lang {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}
