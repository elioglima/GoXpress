package ApiRT

import (
	"encoding/json"
	"goxpress/backend/TransHttp"
	"io/ioutil"
	"net/http"
)

func Info(w http.ResponseWriter, r *http.Request) {
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
	RetornoHttpAPI.HttpStatus = 200
	RetornoHttpAPI.HttpResponse = "Servidor Online"

	js, err := json.Marshal(RetornoHttpAPI)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(js)

}
