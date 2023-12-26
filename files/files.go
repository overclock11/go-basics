package files

import (
	"fmt"
	"goltest/forp"
	"os"
)

var path = "./files/table.txt"

func MultiplicationTable() {
	var table = forp.Multiplications()
	file, err := os.Create(path)

	if err != nil {
		fmt.Println("error al crear el archiivo" + err.Error())
	}
	fmt.Fprintln(file, table)
	file.Close()
}

func AppendTable() {
	var table = forp.Multiplications()
	if !append(path, table) {
		fmt.Println(" fallo agregar datos a la tabla ")
	}
}

func append(path string, table string) bool {
	file, error := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0664)
	if error != nil {
		fmt.Println("falló algo" + error.Error())
		return false
	}

	_, err := file.WriteString(table + "\n")

	if err != nil {
		fmt.Println("fallo" + err.Error())
		return false
	}

	file.Close()
	return true
}

func ReadFile() {
	file, error := os.ReadFile(path)

	if error != nil {
		fmt.Println("fallo otra vez" + error.Error())
	}

	fmt.Println(string(file))
}
