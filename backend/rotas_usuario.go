package backend

import (
	"net/http"
)

func (bc *backend) RTUsuario(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Api Info\n"))

	// db := usuarios.New(global.MyConexao)
	// db.Query("select ativado from usuarios")
	// fmt.Print(w, db.ColCount)
}
