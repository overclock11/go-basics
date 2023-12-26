package variables

import (
	"fmt"
	"strconv"
	"time"
)

// estas variables son visibles en todo el paquete e incluso en los paquetes que importan este paquete
var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	Nombre = "Pedro"
	Estado = true
	Sueldo = 1555.21
	Fecha = time.Now()

	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}

func ConvertToText(number int) (bool, string) {
	var text string = strconv.Itoa(number)
	return true, text
}
