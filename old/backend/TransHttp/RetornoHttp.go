package TransHttp

import (
	"goxpress/backend/global"
	"html/template"
	"time"
)

type retornoTemplate struct {
	KeyAuth          string
	PageTitle        string
	HttpHome         string
	DataProcessameto string
	HttpSolicitado   string
	Token            []byte
	HttpStatus       int
	HttpResponse     string
	HtmlHead         template.HTML
	HtmlScript       template.HTML
	HtmlInputHidden  template.HTML
	HtmlTest         template.HTML
}

func NewRetornoTemplate() *retornoTemplate {

	const layout = "Jan 2, 2006 at 3:04pm (MST)"
	t := time.Now()

	rth := &retornoTemplate{}
	rth.PageTitle = global.PrjTitulo()
	rth.HttpHome = global.HttpHome()
	rth.DataProcessameto = t.UTC().Format(layout)
	rth.HttpSolicitado = global.HttpSolicitado()
	return rth
}
