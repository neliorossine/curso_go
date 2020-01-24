package main

import (
	"encoding/json"
	"fmt"

	"github.com/neliorossine/cursodego/gobuild/model"
)

func main() {

	casa := model.Imovel{}
	casa.Nome = "Casa da Lucy"
	casa.X = 17
	casa.Y = 29
	casa.SetValor(100)

	fmt.Printf("Casa é: %+v\r\n", casa)

	objJSON, _ := json.Marshal(casa)
	fmt.Println("A casa en JSON: ", string(objJSON))

}
