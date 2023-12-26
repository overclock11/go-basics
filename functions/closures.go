package functions

import "fmt"

//devuelve una fgunciona anonima que va devolver un entero
func table(value int) func() int {
	number := value
	secuence := 0

	return func() int {
		secuence++
		return number + secuence
	}
}

func CallClosure() {
	number := 2
	MyTable := table(number)

	for i := 0; i <= 10; i++ {
		fmt.Println(MyTable())
	}
}
