package learnGoWithTests
const (
	spanish = "Spanish"
	french = "French"
	helloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
)

// output hello
func Hello(name string, language string) string{
	if name == "" {
		name = "world"
	}
	return greetPrefix(language) + name
}


func greetPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = helloPrefix
	}
	return
}
