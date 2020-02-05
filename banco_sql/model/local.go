package model

import "database/sql"

//Local aramazena os dados da localidade pelo seu código telefônico
type Local struct {
	Pais             string         `json:"pais" db:"country"`
	Cidade           sql.NullString `json:"cidade" db:"city"`
	CodigoTelefonico int            `json:"cod_telefone" db:"telcode"`
}
