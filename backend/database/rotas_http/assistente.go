package rotas_http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goxpress/backend/TransHttp"
	"goxpress/backend/database/assistente/db_apresentacao"
	"goxpress/backend/database/assistente/db_assistente"
	"goxpress/backend/global"
	"goxpress/backend/libTemplate"
	"goxpress/backend/rotas/publico"
	"html/template"
	"net/http"
	"strconv"
)

type RetornoAssistente struct {
	Id_empresa    int
	Id_assistente int
	Nome          string
	Apresentacao  string
}

func RT_assistente(w http.ResponseWriter, r *http.Request) {

	data := TransHttp.NewRetornoTemplate()
	data.HttpSolicitado = "http://" + r.Host + r.URL.RequestURI()

	buff := new(bytes.Buffer)
	temp, err := libTemplate.Parse("global/HtmlHead")
	temp.Execute(buff, data)
	data.HtmlHead = template.HTML(buff.String())

	buff = new(bytes.Buffer)
	temp, err = libTemplate.Parse("global/HtmlScript")
	temp.Execute(buff, data)
	data.HtmlScript = template.HTML(buff.String())

	buff = new(bytes.Buffer)
	temp, err = libTemplate.Parse("global/HtmlInputHidden")
	temp.Execute(buff, data)
	data.HtmlInputHidden = template.HTML(buff.String())

	temp, err = libTemplate.Parse("inicio")
	if err != nil {
		global.Logger.Erro("falha ao efetuar o parse do .gohtml")
		publico.PageNotFound(w, r)
		return
	}

	id_empresa := 1
	ag := db_assistente.New(global.MyConexao)
	err = ag.GetAll(id_empresa)
	if err != nil {
		fmt.Println(err)
		return
	}

	RA := &RetornoAssistente{}

	for _, assistente := range ag.Rows {

		agp := db_apresentacao.New(global.MyConexao)
		id_assistente, err := strconv.Atoi(assistente[0])

		if err != nil {
			fmt.Println(err)
			return
		}
		agp.GetAll(id_empresa, id_assistente)

		RA.Id_empresa = id_empresa
		RA.Id_assistente = id_assistente
		RA.Nome = assistente[2]
		RA.Apresentacao = agp.Rows[0][3]

	}

	js, err := json.Marshal(RA)
	if err != nil {
		fmt.Println(err)
		return
	}

	data.HtmlTest = template.HTML(js)
	temp.Execute(w, data)

}
