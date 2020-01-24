package model

//Galinha representa uma ave do tipo galinha
type Galinha interface {
	Cacareja() string
}

//Pata representa uma ave do tipo pata
type Pata interface {
	Quack() string
}

//Ave representa um animal
type Ave struct {
	Nome string
}

//Cacareja retorna o som emitido por uma galinha
func (a Ave) Cacareja() string {
	return "CóCóriCó..."
}

//Quack retorna o som emitido por uma pata
func (a Ave) Quack() string {
	return "Quack, quack..."
}
