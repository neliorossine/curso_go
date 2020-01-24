package prefixo

import "strconv"

//Capital representa o número do prefixo de telefone de um estado/província
var Capital = 11

//NomeOperadoraPrefixoInterior nome da operadora e seu prefixo no interior
//var NomeOperadoraPrefixoInterior = operadora.NomeOperadora + " - 13"

var teste = "teste"

//TesteComPrefixo isso é só um teste
var TesteComPrefixo = teste + " " + strconv.Itoa(Capital)
