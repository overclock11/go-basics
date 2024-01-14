package main

import (
	"fmt"
	"goltest/webserver"
	"runtime"
)

// las funciones principales deben llamarse main
func main() {
	//variables.RestoVariables()
	//fmt.Println(variables.ConvertToText(15433))
	//exercises.CoverteAndEvaluate("123")
	//exercises.CoverteAndEvaluate("stirg")
	//getdata.GetFromGekeyboard()
	//fmt.Println(forp.Multiplications())
	//files.AppendTable()
	//files.ReadFile()
	//functions.Exponencial(2)
	//arraysslices.ShowMaps()
	//users.CreateUser()

	/*maria := new(models.Women)
	mario := new(models.Men)
	exercises.TestHumanInterface(maria)
	exercises.TestHumanInterface(mario)*/

	//deferp.APanic()

	// ---------- go routines
	// si termina el codigo y la rutina no ha terminado todo se cierra, hay que poner algo para que espere a que termine la rutina
	// las rutinas de go nos sirven para ejecutar cosas en segundo plano
	/*go goroutines.MyNameDelay("Julian Borray")

	// aqui agrego algo para que se vea la ejecucion de la rutina

	fmt.Println("aca")
	var x string
	fmt.Scanln(&x)*/

	/*channel := make(chan bool)
	goroutines.MyNameDelayWithChannel("Julian", channel)

	defer func() {
		<-channel
	}()*/

	// ---------- go routines

	webserver.MyWebServer()
}

func printMesage(intro string) {
	if intro == "cyberpunk" {
		fmt.Println("bueeeeeenos dias nigth city")
	} else {
		fmt.Println("... musica de metro...")
	}
}

func ConditionalKeywords() {

	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("no es windows, es ", os)
	} else {
		fmt.Println("WIndows a la fija", os)
	}

	// como funciona el switch

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("linux")
	case "osx":
		fmt.Println("osx")
	case "windows":
		fmt.Printf("%s \n\n\n elsiwthc", os)
	default:
		fmt.Printf("%s \n", os)
	}
}
