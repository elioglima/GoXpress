package rotas

import "net/http"

func ApiInfo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Api Info\n"))
}
