package naming

import "fmt"

type Dog struct {
	Name string
	Age  int
}

func (d Dog) Bark(muzzeled bool) {
	if muzzeled {
		fmt.Println("Woof")
	} else {
		fmt.Println("WOOF!!")
	}
}

func Speak(lang string) {
	switch lang {
	case "spanish":
		fmt.Println("Hola")
	default:
		fmt.Println("Hello")
	}
}

func Colour(name string) string {
	switch name {
	case "blue":
		return "#0000FF"
	case "white":
		return "#FFFFFF"
	case "black":
		return "#000000"
	case "grey":
		return "#888888"
	default:
		return "#000000"
	}
}
