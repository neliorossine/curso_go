package model

//BlogPost armazena dados de um posto no Blog
type BlogPost struct {
	UsuarioID int    `json:"userId"`
	ID        int    `json:"id"`
	Titulo    string `json:"title"`
	Texto     string `json:"body"`
}
