package arraysslices

import "fmt"

// declarando unb slice, son elasticos pueden agrandarse en tiempo de ejecución
var tableSlice []int = []int{1, 2, 3, 4, 5}
var arraySlice []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func Slice() {
	// en este caso no muestra ceros en los espacios vacíos, solo muestra los valores que se le asignaron
	fmt.Println(tableSlice)

	partOfArray := arraySlice[3:] // crea desde la posicion 3 hasta el final

	partOfArray = arraySlice[:3] // crea desde el inicio hasta la posicion 3
	fmt.Println(partOfArray)
}

/*fun Capacity() {
	elementsSlyce := make([]int, 5, 20) // crea un slice de 5 posiciones con una capacidad de 20
	fmt.Println(elements, len(elements), cap(elements)) // len() devuelve la longitud del slice, cap() devuelve la capacidad del slice
}*/
