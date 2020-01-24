package main

import (
	"fmt"

	"github.com/neliorossine/cursodego/variaveis/pacotes/operadora"
	"github.com/neliorossine/cursodego/variaveis/pacotes/prefixo"
)

//NomeDoUsuario é o nome do usuário do sistema
var NomeDoUsuario = "Nélio"

func main() {

	fmt.Printf("Nome do usuário : %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo da Capital: %d\r\n", prefixo.Capital)
	fmt.Printf("Nome da operadora: %s\r\n", operadora.NomeOperadora)
	fmt.Printf("Valor de teste: %s\r\n", prefixo.TesteComPrefixo)
}
