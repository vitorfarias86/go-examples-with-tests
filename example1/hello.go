package example1

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Ui, "

/*
Hello world function
*/
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greatingPrefix(language) + name
}
func greatingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
