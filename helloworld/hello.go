package main

import (
	"fmt"
)

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
	arabicHelloPrefix = "Marhaba, "
)

func Hello(name string, language string) string {
	if name == "" {
        name = "world"
    }

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
		case "Spanish":
            prefix = spanishHelloPrefix
		case "French":
			prefix = frenchHelloPrefix
		case "Arabic":
			prefix = arabicHelloPrefix
        default:
            prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Abbas", ""))
}