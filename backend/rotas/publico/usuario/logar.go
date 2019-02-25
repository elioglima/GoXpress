package pub_usuario

import (
	"encoding/json"
	"goxpress/backend/TransHttp"
	"goxpress/backend/rotas/privado"
	"io/ioutil"
	"net/http"
)

func Logar(w http.ResponseWriter, r *http.Request) {
	privado.Direcionar(w, r)
	return

	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	tr := TransHttp.NewRecebeHttp()
	err = json.Unmarshal(body, &tr)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	RetornoHttpAPI := TransHttp.NewRetornoHttpAPI()
	RetornoHttpAPI.HttpStatus = 501
	RetornoHttpAPI.HttpResponse = "Aguardando"

	TokenHttp := TransHttp.NewTokenHttp()
	TokenHttp.Usuario = tr.Usuario
	TokenHttp.Usuario = tr.Senha

	err = TokenHttp.Gerar()
	if err != nil {
		RetornoHttpAPI.HttpResponse = err.Error()

		js, err := json.Marshal(RetornoHttpAPI)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(js)
		return
	}

	// db := usuarios.New(global.MyConexao)
	// db.Query("select ativado  from usuarios")

	// global.Logger.Atencao(db.FieldByName("aivado"))

	RetornoHttpAPI.Token = TransHttp.Token()

	js, err := json.Marshal(RetornoHttpAPI)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	RetornoHttpAPI.HttpStatus = 201
	RetornoHttpAPI.HttpResponse = "Usuário Autenicado com Sucesso"
	w.Write(js)

}
