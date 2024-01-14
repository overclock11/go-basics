package middlwares

import "fmt"

// los middlewares son funciones que se ejecutan antes de la funcion principal del servidor web
// son funciones intermedias que interceptan
// lo que reciben tiene que ser el mismo tipo de dato de lo que devuelven

func sum(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func MyMiddleware() {
	fmt.Println("inicia middleware")

	result := operationsMiddleware(sumar)(2, 3)
}

func operationsMiddleware(function func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("incio el middleware")
		return function(a, b)
	}
}
