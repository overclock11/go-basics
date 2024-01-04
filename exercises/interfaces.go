package exercises

import (
	"fmt"
	"goltest/interfaces"
)

func TestHumanInterface(human interfaces.Human) {
	human.Eat()
	human.Sleep()
	human.Work()
	human.IsAdult()
	fmt.Println("cargo ")
}
