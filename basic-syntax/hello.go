package basic

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
func SayHello(name string, language string) string {
	if name == "" {
		name = "Go with test"
	}

	var prefix string
	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix + name
}