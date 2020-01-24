package main

import (
	"fmt"

	"github.com/neliorossine/cursodego/variaveis/funcoes_basico/matematica"
)

func main() {
	//resultado := multiplicador(2, 2)

	resultado := matematica.Calculo(matematica.Multiplicador, 2, 2)
	fmt.Printf("O resultado da multiplicação foi: %d\r\n", resultado)

	resultado = matematica.Soma(3, 3)
	fmt.Printf("O resultado da soma foi: %d\r\n", resultado)

	resultado = matematica.Calculo(matematica.Divisor, 12, 3)
	fmt.Printf("O resultado da divisão foi: %d\r\n", resultado)

	resultado, resto := matematica.DivisorComResto(12, 5)
	fmt.Printf("O resultado da divisão foi: %d e o resto da divisão foi %d\r\n", resultado, resto)
}

/*func multiplicador(x int, y int) int {
	return x * y
}*/
