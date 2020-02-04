package model

//Local aramazena os dados da localidade pelo seu código telefônico
type Local struct {
	País             string `json:"pais" db:"country" bson:"country"`
	Cidade           string `json:"cidade" db:"city" bson:"city"`
	CodigoTelefonico int    `json:"cod_telefone" db:"telcode" bson:"telcode"`
}
