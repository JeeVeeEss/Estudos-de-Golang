package main

import (
	"fmt"
	"unsafe"
)

var pi int = 3

func main() {
	var alt *int = &pi
	fmt.Printf("%p : %d\n", &pi, *alt)
	fmt.Printf("Tamanho da variável: %d\nTamanho do ponteiro: %d", unsafe.Sizeof(pi), unsafe.Sizeof(alt))
}
