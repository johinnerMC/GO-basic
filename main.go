package main

import "fmt"

func CalcularVolumenCaja(largo, ancho, alto int) (a int, d string) {
	texto := "litros"
	return largo * ancho * alto, texto
}

func main() {
	value1, value2 := CalcularVolumenCaja(50, 50, 30)
	divicion := value1 / 3
	fmt.Println("la capacidad de la caja en litros es:", divicion, value2)

	fmt.Printf("La capacidad de la caja en litros es: %d %s", divicion, value2)
}
