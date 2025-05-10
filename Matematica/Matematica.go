/*
Objetivo: Criar funções para aplicar
em situações envolvendo Matemática.
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var Vx int32
	var Vy int32
	fmt.Printf("Digite o Coeficiente X: ")
	fmt.Scan(&Vx)
	fmt.Printf("Digite o Coeficiente Y: ")
	fmt.Scan(&Vy)

	Normalizar_vetor(Vx, Vy)

}

// Função para calcular a norma de um vetor bidimensional.
func Normalizar_vetor(Vx int32, Vy int32) int32 {
	norma := math.Pow(math.Pow(float64(Vx), 2)+math.Pow(float64(Vy), 2), 0.5)
	fmt.Printf("Norma do vetor: %d", int32(norma))
	return int32(norma)
}
