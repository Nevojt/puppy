package puppy

import "github.com/Nevojt/dog"

func Burk() string {
	return "Woof!"
}

func Burks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrowUp(BigBark())

}

func BigBarks() string {
	return dog.WhenGrowUp(BigBarks())
}
