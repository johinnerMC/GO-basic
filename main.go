package main

import "fmt"

func main() {
	// declacion de variables
	helloMessage := "hello"
	worldMessage := "world"

	//	Println
	fmt.Println(helloMessage)
	fmt.Println(worldMessage)

	// Printf
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos)

	// Sprintf
	messege := fmt.Sprintf("%s tiene mas de %d cursos", nombre, cursos)
	fmt.Println(messege)

	// 	Tipo datos
	fmt.Printf("nombre: %T\n", nombre)
	fmt.Printf("cursos: %T\n", cursos)

}
