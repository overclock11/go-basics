package models

// herencia
type Women struct {
	Men
}

func (women Women) Eat() {
	women.eat = true
}

func (women Women) Sleep() {
	women.sleep = true
}

func (women Women) Work() {
	women.work = false
}

func (women Women) IsAdult() (isAdult bool) {
	return true
}
