package arraysslices

import "fmt"

func ShowMaps() {
	// entre corchetes se pone el tipo de dato de la clave
	countries := make(map[string]string)

	countries["MX"] = "Mexico"
	countries["CO"] = "Colombia"
	countries["AR"] = "Argentina"
	countries["BR"] = "Brasil"
	fmt.Println(countries)
	fmt.Println(countries["CO"])

	championship := map[string]int{
		"Barcelona":       39,
		"Real Madrid":     38,
		"Bayern Munich":   7,
		"Atletico Madrid": 7,
	}

	fmt.Println(championship)

	for team, score := range championship { //range es como un foreach
		fmt.Printf("The team %s has %d points\n", team, score)
	}

	//elimiando elementos de un mapa
	delete(championship, "Barcelona")

	if score, exist := championship["Atletico Madrid"]; exist {
		fmt.Printf("has %d points\n", score)
	} else {
		fmt.Println("El equipo no se encuentra en el campeonato")
	}

}
