package main

import "fmt"

const englishHelloPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	return getPrefix(lang) + name
}

func getPrefix(lang string) (prefix string) {
	switch lang {
	case "French":
		prefix = frenchHelloPrefix
	case "English":
		prefix = englishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}
func main() {
	fmt.Println(Hello("Kyo", "English"))
}
