package helloworld

const messagePrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return messagePrefix + name
}
