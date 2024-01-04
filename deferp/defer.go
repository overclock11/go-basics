package deferp

import "fmt"

// ultima instrucción en ejecutarse al salir de una función, pueden ser  varios defer
func deferTest() {
	fmt.Println("inicio")
	defer fmt.Println("fin") // se ejecuta al final siempre
	fmt.Println("medio")
}

// interrumpe la ejecución del sistema
func APanic() {
	defer fmt.Println("fin")
	panic("panic")
}
