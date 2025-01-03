package puppy

import (
	"fmt"
	"github.com/Nevojt/dog"
)

//const name = "Bosk"

func Burk() string {
	return "Woof!"
}

func Burks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrowUp(Burk())

}

func BigBarks() string {
	return dog.WhenGrowUp(Burks())
}

func From11() {
	fmt.Println("I'm version 1.1.0")
}

func From12() {
	fmt.Println("I'm version 1.2.0")
}

func From13() {
	fmt.Println("I'm version 1.3.0")
}

func GoodDog(s string) string {
	return dog.GoodBoy(s)
}

// Tags
