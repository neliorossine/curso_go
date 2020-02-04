package manipulador

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/neliorossine/cursodego/banco_sql/model"
	"github.com/neliorossine/cursodego/banco_sql/repo"
)

//Local é o manipulador da requisição de rota /local/
func Local(w http.ResponseWriter, r *http.Request) {

	local := model.Local{}
	codigoTelefone, err := strconv.Atoi(r.URL.Path[7:])
	if err != nil {
		http.Error(w, "Não foi enviado um número válido. Verifque.", http.StatusBadRequest)
		fmt.Println("[local] Erro ao converter o número enviado: ", err.Error())
		return
	}

	sql := "selct country, city, telcode from place where telcode = ?"
	linha, err := repo.Db.Queryx(sql, codigoTelefone)
	if err != nil {
		http.Error(w, "Não foi possível pesquisar esse número.", http.StatusInternalServerError)
		fmt.Println("[local] não foi possível executar a query: ", sql, " Erro: ", err.Error())
		return
	}

	for linha.Next() {
		err = linha.StructScan(&local)
		if err != nil {
			http.Error(w, "Não foi possível pesquisar esse número.", http.StatusInternalServerError)
			fmt.Println("[local] não foi possível fazer o binding dos dados do banco na struct: ", err.Error())
			return
		}
	}

	if err := ModeloLocal.ExecuteTemplate(w, "local.html", local); err != nil {
		http.Error(w, "Houve um erro na renderização da página.", http.StatusInternalServerError)
		fmt.Println("[local] Erro na execucao do modelo: ", err.Error())
	}

	sql = "insert into logquery (daterequest) values (?)"
	resultado, err := repo.Db.Exec(sql, time.Now().Format("02/01/2006 15:04:05"))
	if err != nil {
		fmt.Println("[local] Erro na inclusão do log: ", sql, " - ", err.Error())
	}
	linhasAfetadas, err := resultado.RowsAffected()
	if err != nil {
		fmt.Println("[local] Erro ao pegar o número de linhas afetadas pelo comando: ", sql, " - ", err.Error())
	}
	fmt.Println("Sucesso!", linhasAfetadas, " linha(s) afetada(s)")

}
