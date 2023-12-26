package functions

import "fmt"

func Calcs() {

	var number3 = 3

	//funcion anonima asignada a una variable
	sum := func(number1 int, number2 int) int {
		return number1 + number2*number3
	}

	fmt.Println(sum(2, 4))

}
