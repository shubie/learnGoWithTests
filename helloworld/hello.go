package main

import (
	"fmt"
)

const (
	english = "english"
	french  = "french"
	spanish = "spanish"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Boujor, "
)

func HelloWorld(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return getLanguagePrefix(language) + name
}

func getLanguagePrefix(language string) (prefix string) {
	switch language {
	case english:
		prefix = englishHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	}
	return prefix
}

func main() {
	fmt.Println(HelloWorld("Shuaib", "spanish"))
}
