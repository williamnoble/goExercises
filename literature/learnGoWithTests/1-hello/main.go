package main

import "fmt"

const helloPrefix = "Hello, "
const helloPrefixSpanish = "Hola, "
const helloPrefixFrench = "Bonjour, "

func Hello(name, language string) string {

	if name == "" {
		name = "World"
	}

	prefix := generatePrefix(language)

	return fmt.Sprintf("%s%s", prefix, name)
}

// Note by specifying the prefix string in signature
// it makes this variable available throughout the function
func generatePrefix(language string) (prefix string) {
	// var prefix string
	switch language {
	case "French":
		prefix = helloPrefixFrench
	case "Spanish":
		prefix = helloPrefixSpanish
	default:
		prefix = helloPrefix
	}
	return prefix
}

func main() {
	name := "William"
	fmt.Println(Hello(name, "Spanish"))
}
