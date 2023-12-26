package getdata

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func GetFromGekeyboard() {
	var number1 int
	var number2 int
	var label string
	var error error

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("ingrese numero 1")
	if scanner.Scan() {
		number1, error = strconv.Atoi(scanner.Text())
		if error != nil {
			panic("se jodio")
		}
	}

	fmt.Println("ingrese numero 2")
	if scanner.Scan() {
		number2, error = strconv.Atoi(scanner.Text())
		if error != nil {
			panic("se jodio en el nummero 2")
		}
	}

	fmt.Println("un texto string")
	if scanner.Scan() {
		label = scanner.Text()
	}

	fmt.Println(label, number1*number2)
}
