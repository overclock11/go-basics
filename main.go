package main

import (
	"fmt"
	arraysslices "goltest/arrays_slices"
	"goltest/functions"
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
	functions.Exponencial(2)
	arraysslices.ShowArrays()
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
