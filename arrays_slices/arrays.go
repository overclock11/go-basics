package arraysslices

import "fmt"

// definicio de un array, tiene que tener definida la cantidad de posiciones
// si no la tiene definida se convierte en un slice
// todas las posiciones se inicializan en 0
var table [10]int = [10]int{10, 3, 312, 3123}

func ShowArrays() {
	table[2] = 231
	table[9] = 96727
	fmt.Println(table)

	for i := 0; i < len(table); i++ {
		fmt.Println(table[i])
	}
}
