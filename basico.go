package main

import (
	"fmt"

	"matematica.com/data/GA"
)

// const pi float32 = 3.1415
var (
	msg    string
	letter []string
	//pedaco []string
)

func main() {

	fmt.Printf("Mensagem: ")
	fmt.Scanln(&msg) /* Receber um valor de entrada e assinar o valor no endereço da variável msg (&msg) */
	fmt.Printf("Mensagem recebida: %s\n", msg)
	fmt.Println(&msg)
	divider()
	//letters := make([]string, 0, len(msg))
	for i := 0; i < len(msg); i++ {
		fmt.Printf("%b : %c \n", msg[i], msg[i])
		letter = append(letter, fmt.Sprintf("%b", msg[i]))
		fmt.Println(letter)
	}
	divider()
	a := test(12)
	fmt.Println(a)
	/* for i := 0; i <= len(msg); i++ {
		fmt.Println(string(msg[i]))
	} */
	divider()
	fmt.Println("Geometria analítica iniciada:")
	GA.Normalizar_vetor(12, 21)
	divider()
	fmt.Println(letter)
}

func test(valor int32) int32 {
	fmt.Printf("WORKING -- ")
	return int32(valor * valor)
}
func divider() {
	fmt.Println("**********")
}
