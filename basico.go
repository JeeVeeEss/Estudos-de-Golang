package main

import (
	"fmt"

	"matematica.com/data/GA"
)

// const pi float32 = 3.1415
var (
	msg string
	//pedaco []string
)

func main() {

	fmt.Printf("Mensagem: ")
	fmt.Scanln(&msg) /* Receber um valor de entrada e assinar o valor no endereço da variável msg (&msg) */
	fmt.Printf("Mensagem recebida: %s\n", msg)
	fmt.Println(&msg)
	a := test(12)
	fmt.Println(a)
	/* for i := 0; i <= len(msg); i++ {
		fmt.Println(string(msg[i]))
	} */
	fmt.Println("Geometria analítica iniciada:")
	GA.Normalizar_vetor(12, 21)
}

func test(valor int32) int32 {
	fmt.Printf("WORKING -- ")
	return int32(valor * valor)
}
