package exercises

import (
	"fmt"
	"strconv"
)

// para ser pública arrancan con mayuscula
func CoverteAndEvaluate(stringNumber string) (int, string) {
	// con := crea la variable, asume que lo que se pone antes no existe

	number, error := strconv.Atoi(stringNumber)
	// number, _ := strconv.Atoi(stringNumber) para descartar la segunda variable que retorna, la de error se puede usar _
	if error != nil {
		return 0, "mayor a 100"
	}
	fmt.Println(error)
	if number > 100 {
		return number, "mayor a 100"
	} else {
		return number, "menor a 100"
	}
}
