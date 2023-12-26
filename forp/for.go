package forp

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Multiplications() string {
	number := 0
	scanner := bufio.NewScanner(os.Stdin)
	tabla := ""

	fmt.Println("Ingrese el numero de la tabla de multiplicar")
	if scanner.Scan() {
		number, _ = strconv.Atoi(scanner.Text())
		for i := 0; i < 10; i++ {
			tabla += fmt.Sprintf("%d x %d = %d \n", number, i, number*i)
		}
	}
	return tabla
}
