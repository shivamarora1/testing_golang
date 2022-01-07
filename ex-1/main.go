package main

import "fmt"

const GREETING_ENGLISH = "Hey "
const GREETING_SPANISH = "Hola "
const GREETING_FRENCH = "Bonjour "

const LANGUAGE_ENGLISH = "english"
const LANGUAGE_SPANISH = "spanish"
const LANGUAGE_FRENCH = "french"

func helloPrefix(language string) string {
	switch language {
	case LANGUAGE_ENGLISH:
		return GREETING_ENGLISH
	case LANGUAGE_SPANISH:
		return GREETING_SPANISH
	case LANGUAGE_FRENCH:
		return GREETING_FRENCH
	default:
		return GREETING_ENGLISH
	}
}

func sayHello(language string, name string) string {
	prefix := helloPrefix(language)
	if name != "" {
		return prefix + name
	} else {
		return ""
	}
}

func main() {
	fmt.Printf("%v", sayHello(LANGUAGE_ENGLISH, "Nick"))
}
