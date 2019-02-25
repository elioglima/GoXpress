package publico

import (
	"goxpress/backend/TransHttp"
	"goxpress/backend/global"
	"goxpress/backend/libTemplate"
	"net/http"
	"strconv"
)

func PageNotFound(w http.ResponseWriter, r *http.Request) {

	data := TransHttp.NewRetornoTemplate()
	data.HttpSolicitado = "http://" + r.Host + r.URL.RequestURI()

	temp, err := libTemplate.Parse("PageNotFound")
	if err != nil {
		global.Logger.Erro("Erro")
	}
	temp.Execute(w, data)

}

func ErroIntenoServidor(w http.ResponseWriter, r *http.Request) {
	data := TransHttp.NewRetornoTemplate()
	data.HttpSolicitado = "http://" + r.Host + r.URL.RequestURI()
	i, err := strconv.Atoi(r.Header.Get("HttpStatus"))

	if err != nil {
		data.HttpStatus = 0
	} else {
		data.HttpStatus = i
	}

	data.HttpResponse = r.Header.Get("HttpResponse")

	temp, err := libTemplate.Parse("ErroIntenoServidor")
	if err != nil {
		global.Logger.Erro("Erro")
	}
	temp.Execute(w, data)

}
