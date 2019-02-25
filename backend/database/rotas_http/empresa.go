package rotas_http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goxpress/backend/TransHttp"
	"goxpress/backend/database/empresa/db_empresa"
	"goxpress/backend/global"
	"goxpress/backend/libTemplate"
	"goxpress/backend/rotas/publico"
	"html/template"
	"net/http"
)

func RT_empresa(w http.ResponseWriter, r *http.Request) {

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

	emp := db_empresa.New(global.MyConexao)
	err = emp.GetID(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	js, err := json.Marshal(emp.Rows[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	data.HtmlTest = template.HTML(js)
	temp.Execute(w, data)

}
