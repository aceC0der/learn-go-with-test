package basic

const englishHelloPrefix = "Hello, "
func SayHello(name string) string {
	if name=="" {
		name = "Go with test"
	}
	return englishHelloPrefix + name
}