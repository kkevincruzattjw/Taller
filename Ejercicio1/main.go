package main

import "fmt"

func main() {
	var palabra string
	var cantidad int

	conteo := make(map[string]int)

	fmt.Println("Ingrese 5 palabras:")

	for i := 0; i < 5; i++ {
		fmt.Scan(&palabra)
		conteo[palabra]++
	}

	fmt.Println("Conteo de palabras:")

	palabraMayor := ""
	cantidadMayor := 0

	for palabra, cantidad = range conteo {
		fmt.Println(palabra, ":", cantidad)

		if cantidad > cantidadMayor {
			cantidadMayor = cantidad
			palabraMayor = palabra
		}
	}

	fmt.Println("La palabra que más se repite es:", palabraMayor)
	fmt.Println("Número de repeticiones:", cantidadMayor)
}
