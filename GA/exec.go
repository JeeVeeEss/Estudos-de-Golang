/*
Objetivo: Criar funções para aplicar
em situações envolvendo Matemática.
*/
package GA

import (
	"fmt"
	"math"
)

type Vetor_R2 struct {
	value_i float32
	value_j float32
}

// Vetor Tridimensional
type Vetor_R3 struct {
	value_i float32
	value_j float32
	value_k float32
}

// Função para calcular a norma de um vetor bidimensional.
func Normalizar_vetor(Vx int32, Vy int32) int32 {
	norma := math.Pow(math.Pow(float64(Vx), 2)+math.Pow(float64(Vy), 2), 0.5)
	fmt.Printf("Norma do vetor: %d", int32(norma))
	return int32(norma)
}
