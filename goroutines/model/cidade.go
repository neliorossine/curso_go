package model

//Cidade represente a cidade e estado do Brasil
type Cidade struct {
	Nome   string `json:"nome"`
	Estado string `json:"estado"`
}
