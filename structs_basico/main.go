package main

import "fmt"

//Imovel é uma struct que armazena dados de um imovel
type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func main() {

	casa := Imovel{}
	fmt.Printf("A casa é: %+v\r\n", casa)

	apartamento := Imovel{17, 56, "Meu AP", 760000}
	fmt.Printf("O Apartamento é: %+v\r\n", apartamento)

	chacara := Imovel{
		Y:     85,
		Nome:  "Chácara",
		valor: 55,
		X:     22,
	}
	fmt.Printf("A chácara é: %+v\r\n", chacara)

	casa.Nome = "Lar Doce Lar"
	casa.valor = 450000
	casa.X = 18
	casa.Y = 31
	fmt.Printf("A casa é: %+v\r\n", casa)

	casaCampo := new(Imovel)
	fmt.Printf("Casa de campo é : %p - %+v\r\n", &casaCampo, casaCampo)

	chacaraCampo := Imovel{17, 28, "Chácara Linda", 2800000}
	fmt.Printf("Chácara de campo é : %p - %+v\r\n", &chacaraCampo, chacaraCampo)

	mudaImovel(&chacaraCampo)
	fmt.Printf("Chácara de campo é : %p - %+v\r\n", &chacaraCampo, chacaraCampo)

}

func mudaImovel(imovel *Imovel) {
	imovel.X = imovel.X + 10
	imovel.Y = imovel.Y - 5
}
