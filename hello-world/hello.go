package main

func SayGreetingsTo(name string, language Language) string {
	prefix := GetGreetingsPrefix(language)

	if name == "" {
		name = "World"
	}

	return prefix + ", " + name + "!"
}

func GetGreetingsPrefix(language Language) string {
	switch language {
	case Spanish:
		return "Hola"
	case French:
		return "Bonjour"
	}

	return "Hello"
}
