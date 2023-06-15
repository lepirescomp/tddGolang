package TDDPractice

import "fmt"

const spanishGretting = "Hola"
const portugueseGretting = "Olá"
const englishGretting = "Hello"

func greeting(language string) (g string) {
	switch language {
	case "spanish":
		return spanishGretting
	case "portuguese":
		return portugueseGretting
	default:
		return englishGretting
	}
}

func Hello(language string, name string) string {
	g := greeting(language)
	if name != "" {
		return g + ", " + name
	}
	return name
}

func NestedLoop() int {
	var sum int
	for i := 0; i < 100; i += 1 {
		for j := 0; j < 100; j += 1 {
			sum += i + j
		}
	}
	return sum
}

func hello() {
	s := Hello("english", "Test")
	fmt.Print(s)
}
