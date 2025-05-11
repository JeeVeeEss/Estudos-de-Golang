/*
Objetivo: Criar funções para aplicar
em situações envolvendo Matemática.
*/
package GA

import (
	"fmt"
	"math"
)

// Função para calcular a norma de um vetor bidimensional.
func Normalizar_vetor(Vx int32, Vy int32) int32 {
	norma := math.Pow(math.Pow(float64(Vx), 2)+math.Pow(float64(Vy), 2), 0.5)
	fmt.Printf("Norma do vetor: %d", int32(norma))
	return int32(norma)
}
