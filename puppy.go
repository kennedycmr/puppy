package puppy

import (
	"github.com/kennedycmr/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() {
	return dog.WhenGrownUp(Barks())
}
