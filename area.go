package main

import "fmt"

const pi float32 = 3.1415

func main() {
	var raio float32
	fmt.Printf("Insira o raio: ")
	fmt.Scan(&raio)

	var area float32 = pi * (raio * raio)

	fmt.Printf("A área é de: %f", area)

}
