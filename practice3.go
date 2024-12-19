package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	count := 0
	var nombre string
	var edad int

	userlist := [10]Person{}
	for {
		var back string
		fmt.Println("B(Back) C(Continue)")
		fmt.Println("B/C")
		fmt.Scan(&back)
		if back == "B" || back == "b" {
			count--
		} else if back == "C" || back == "c" {
			fmt.Println("...")
		}
		fmt.Println("-- User Creator --")
		fmt.Println("--Ingresa el Nombre de la Persona--")
		fmt.Scan(&nombre)
		fmt.Println("Ingresa la Edad")
		fmt.Scan(&edad)
		fmt.Println("Tu Nombre es: " + nombre)
		user := Person{
			Name: nombre,
			Age:  edad,
		}
		fmt.Println("Persona Creada Con Exito: ", user)
		var cond string
		fmt.Println("¿Deseas Añadirlo a la lista? Y/N")
		fmt.Scan(&cond)
		switch cond {
		case "Y", "y":
			if count >= 0 {
				userlist[count] = user
			} else {
				fmt.Println("El indice es menor a 0...")
				count++
			}

			fmt.Println("Listo, lo as añadido a la lista: ", userlist)
			fmt.Println("Agregado En Indice: ", count)
			count++
		case "N", "n":
			fmt.Println("Ok...")
		case "B", "b":
			count--
		}
		if count > 10 {
			break
		}
	}

}
