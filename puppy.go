package puppy

import (
	"github.com/kevinrosenfield/dog"
)

func Bark() (string, string) {
	dogBark := dog.Bark()
	return "Woof", dogBark
}

func Barks() string {
	dogBarks := dog.Barks()
	return dogBarks
}
