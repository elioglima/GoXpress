package rotas_http

import (
	"encoding/json"
	"goxpress/backend/TransHttp"
	"goxpress/backend/database/assistente/db_resposta_acao"
	"goxpress/backend/global"
	"io/ioutil"
	"net/http"
)

func RT_resposta_acao_add(w http.ResponseWriter, r *http.Request) {

	RetornoHttpAPI := TransHttp.NewRetornoHttpAPI()
	RetornoHttpAPI.HttpStatus = 0
	RetornoHttpAPI.HttpResponse = "Aguardando Mensagem"

	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	Request := db_resposta_acao.CRow{}
	err = json.Unmarshal(body, &Request)
	if err != nil {
		RetornoHttpAPI.HttpStatus = 501
		RetornoHttpAPI.HttpResponse = err.Error()

	} else if Request.Id_empresa == 0 {
		RetornoHttpAPI.HttpStatus = 501
		RetornoHttpAPI.HttpResponse = "Id empresa não informado"

	} else if Request.Id_assistente == 0 {
		RetornoHttpAPI.HttpStatus = 501
		RetornoHttpAPI.HttpResponse = "Id assistente não informado"

	} else if len(Request.Descricao) == 0 {
		RetornoHttpAPI.HttpStatus = 501
		RetornoHttpAPI.HttpResponse = "Descrição não informada"

	}

	if RetornoHttpAPI.HttpStatus == 0 {
		RSPA := db_resposta_acao.New(global.MyConexao)
		err = RSPA.Add(Request)
		if err != nil {
			RetornoHttpAPI.HttpStatus = 201
			RetornoHttpAPI.HttpResponse = err.Error()
		}
	}

	js, err := json.Marshal(RetornoHttpAPI)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(js)

}

func RT_resposta_acao_edd(w http.ResponseWriter, r *http.Request) {

	RetornoHttpAPI := TransHttp.NewRetornoHttpAPI()
	RetornoHttpAPI.HttpStatus = 0
	RetornoHttpAPI.HttpResponse = "Aguardando Mensagem"

	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	Request := db_resposta_acao.CRow{}
	err = json.Unmarshal(body, &Request)
	if err != nil {
		RetornoHttpAPI.HttpStatus = 501
		RetornoHttpAPI.HttpResponse = err.Error()

	} else if Request.Id == 0 {
		RetornoHttpAPI.HttpStatus = 501
		RetornoHttpAPI.HttpResponse = "Id de pesquisa não informado"

	} else if Request.Id_empresa == 0 {
		RetornoHttpAPI.HttpStatus = 501
		RetornoHttpAPI.HttpResponse = "Id empresa não informado"

	} else if Request.Id_assistente == 0 {
		RetornoHttpAPI.HttpStatus = 501
		RetornoHttpAPI.HttpResponse = "Id assistente não informado"

	} else if len(Request.Descricao) == 0 {
		RetornoHttpAPI.HttpStatus = 501
		RetornoHttpAPI.HttpResponse = "Descrição não informada"

	}

	if RetornoHttpAPI.HttpStatus == 0 {
		RSPA := db_resposta_acao.New(global.MyConexao)
		err = RSPA.Edd(Request)
		if err != nil {
			RetornoHttpAPI.HttpStatus = 501
			RetornoHttpAPI.HttpResponse = err.Error()
		} else {
			RetornoHttpAPI.HttpStatus = 201
			RetornoHttpAPI.HttpResponse = "Dados atualizado com sucesso"
		}
	}

	js, err := json.Marshal(RetornoHttpAPI)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(js)

}
