package backend

import (
	"goxpress/backend/global"
	"net/http"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"strings"

	"goxpress/libs"

	"github.com/gorilla/mux"
)

type backend struct {
	logger   *libs.Logs
	Arquivos *CnfArq
	Rotas    *mux.Router
	DataView *SDataView
}

func NewBackEnd() *backend {
	return &backend{}
}

type customFileServer struct {
	root            http.Dir
	NotFoundHandler func(http.ResponseWriter, *http.Request)
}

func (fs *customFileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//if empty, set current directory
	dir := string(fs.root)
	if dir == "" {
		dir = "."
	}

	//add prefix and clean
	upath := r.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		r.URL.Path = upath
	}
	upath = path.Clean(upath)

	//path to file
	name := path.Join(dir, filepath.FromSlash(upath))

	//check if file exists
	f, err := os.Open(name)
	if err != nil {
		if os.IsNotExist(err) {
			fs.NotFoundHandler(w, r)
			return
		}
	}
	defer f.Close()

	http.ServeFile(w, r, name)
}

func CustomFileServer(root http.Dir, NotFoundHandler http.HandlerFunc) http.Handler {
	return &customFileServer{root: root, NotFoundHandler: NotFoundHandler}
}

func (bc *backend) Ini() {

	global.Load()
	// inicio do programa ou servidor

	global.Logger.Sucesso("Iniciando Servidor ...")

	// sDir := "GoXpress"
	// ftpClient.Conectar()
	// ftpClient.MkDir(sDir)
	// ftpClient.ChangeDir(sDir)
	// ftpClient.Fechar()

	global.Logger.Atencao("Conectando ao banco de dados ...")

	global.ConctarDB()
	global.Logger.Atencao("Serviço de Http")

	bc.Configs()
	bc.setRotas()

	sPorta := "2525"
	go func() {
		err := http.ListenAndServe(":"+sPorta, bc.Rotas)
		if err != nil {
			global.Logger.Erro("ListenAndServe: ", err)
		}
	}()

	global.Logger.Sucesso("Servidor Iniciado com sucesso Http: " + sPorta + " ...")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	global.Logger.Sucesso("Finalizando servidor")
	os.Exit(0)

}
