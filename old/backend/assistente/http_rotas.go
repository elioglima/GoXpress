package assistente

import (
	"encoding/json"
	classe "goxpress/backend/Classes"
	"goxpress/backend/TransHttp"
	"io/ioutil"
	"net/http"
)

func ResponseWriterJson(w http.ResponseWriter, RetornoHttpAPI *TransHttp.RetornoHttpAPI) error {
	js, err := json.Marshal(RetornoHttpAPI)
	if err != nil {
		w.Write([]byte(err.Error()))
		return err
	}

	w.Write(js)
	return nil
}

func RT_assistente(w http.ResponseWriter, r *http.Request) {

	RetornoHttpAPI := TransHttp.NewRetornoHttpAPI()
	RetornoHttpAPI.HttpStatus = 0
	RetornoHttpAPI.HttpResponse = "Autenticando Solicitação"

	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	Identificador := &classe.CAgenteIdentificador{}
	err = json.Unmarshal(body, Identificador)
	if err != nil {
		RetornoHttpAPI.HttpStatus = 501
		RetornoHttpAPI.HttpResponse = err.Error()
		ResponseWriterJson(w, RetornoHttpAPI)
		return

	} else if Identificador.Tecnologia == 0 {
		RetornoHttpAPI.HttpStatus = 501
		RetornoHttpAPI.HttpResponse = "Tipo de Autenticação não Informado"
		ResponseWriterJson(w, RetornoHttpAPI)
		return

	} else if len(Identificador.Identificador) == 0 {
		RetornoHttpAPI.HttpStatus = 501
		RetornoHttpAPI.HttpResponse = "Identificador não informado"
		ResponseWriterJson(w, RetornoHttpAPI)
		return

	} else if len(Identificador.Mensagem) == 0 {
		RetornoHttpAPI.HttpStatus = 501
		RetornoHttpAPI.HttpResponse = "Mensagem não informada"
		ResponseWriterJson(w, RetornoHttpAPI)
		return

	}

	RetornoHttpAPI.HttpStatus = 201
	RetornoHttpAPI.HttpResponse = "Identificador Autenticado com Sucesso"

	go func() {
		Agente := NewCAgente()
		Agente.Start(Identificador)
		defer Agente.Close()
	}()

	ResponseWriterJson(w, RetornoHttpAPI)
}
