package model

import "errors"

//Imovel representa informações de um imóvel
type Imovel struct {
	Nome  string `json:"nome"`
	valor int
	X     int `json:"coordenada_x"`
	Y     int `json:"coordeada_y"`
}

/*
ErrValorDeImovelInvalido é um erro que é apresentado quando é atribuído a um imóvel com valor muito baixo
*/
var ErrValorDeImovelInvalido = errors.New("O valor informado não é válido para um imóvel")

/*
ErrValorDeImovelMuitoAlto é um erro que é apresentado quando é atribuído ao imóvel um valor muito alto, fora do mercado.
*/
var ErrValorDeImovelMuitoAlto = errors.New("O valor informado é muito alto")

//GetValor obtem o valor do Imovel
func (i *Imovel) GetValor() int {
	return i.valor
}

//SetValor define o valor do imovel
func (i *Imovel) SetValor(valor int) (err error) {
	err = nil
	if valor < 50000 {
		err = ErrValorDeImovelInvalido
		return
	} else if valor > 10000000 {
		err = ErrValorDeImovelMuitoAlto
		return
	}

	i.valor = valor
	return
}
