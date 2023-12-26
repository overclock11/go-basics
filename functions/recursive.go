package functions

import "fmt"

func Exponencial(value int) {
	if value > 20000 {
		return
	}
	fmt.Println(value)
	Exponencial(value * 2)
}
