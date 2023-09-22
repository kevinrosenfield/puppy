package puppy

import (
	"fmt"

	"github.com/kevinrosenfield/dog"
)

func Bark() {
	dogBark := dog.Bark()

	fmt.Println(dogBark)
}
