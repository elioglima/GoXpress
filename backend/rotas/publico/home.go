package publico

import (
	"bytes"
	"goxpress/backend/TransHttp"
	"goxpress/backend/database/empresa/db_empresa"
	"goxpress/backend/global"
	"goxpress/backend/libTemplate"
	"html/template"
	"net/http"
	"reflect"
)

func getStructTag(f reflect.StructField) string {
	return string(f.Tag)
}

func Home(w http.ResponseWriter, r *http.Request) {

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
		PageNotFound(w, r)
		return
	}

	emp := db_empresa.New(global.MyConexao)
	global.Logger.Sucesso(emp.Nome_Tabela + " - ok")
	err = emp.GetID(1)

	if err != nil {
		global.Logger.Erro("falha ao efetuar o parse do .gohtml")
		PageNotFound(w, r)
		return
	}

	for _, row := range emp.Rows {
		global.Logger.Sucesso(row[1])
	}

	data.HtmlTest = template.HTML(" ")
	temp.Execute(w, data)

}
