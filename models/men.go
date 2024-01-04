package models

type Men struct {
	age    int
	height int
	weight int
	eat    bool
	sleep  bool
	work   bool
}

func (men Men) Eat() {
	men.eat = true
}

func (men Men) Sleep() {
	men.sleep = true
}

func (men Men) Work() {
	men.work = false
}

func (men Men) IsAdult() bool {
	return true
}
