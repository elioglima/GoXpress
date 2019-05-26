package privado

import (
	"bytes"
	"goxpress/backend/TransHttp"
	"goxpress/backend/global"
	"goxpress/backend/libTemplate"
	"goxpress/backend/rotas/publico"
	"html/template"
	"net/http"
)

func Direcionar(w http.ResponseWriter, r *http.Request) {

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

	temp, err = libTemplate.Parse("PageRedirect")
	if err != nil {
		global.Logger.Erro("falha ao efetuar o parse do .gohtml")
		publico.PageNotFound(w, r)
		return
	}

	temp.Execute(w, data)

}
