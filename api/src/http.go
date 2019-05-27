package src

import (
	"net/http"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"strings"

	"GoXpress/api/src/global"
	"GoXpress/api/src/libs"

	"github.com/gorilla/mux"
)

type api struct {
	logger *libs.Logs
	Rotas  *mux.Router
}

func NewApi() *api {
	return &api{}
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

func (api *api) Ini() {

	global.Load()
	// inicio do programa ou servidor

	global.Logger.Println("Iniciando Servidor ...")

	// sDir := "GoXpress"
	// ftpClient.Conectar()
	// ftpClient.MkDir(sDir)
	// ftpClient.ChangeDir(sDir)
	// ftpClient.Fechar()

	// global.Logger.Println("Conectando ao banco de dados ...")

	// global.ConctarDB()
	// global.Logger.Atencao("Serviço de Http")

	// api.Configs()
	api.setRotas()

	sPorta := "2525"
	go func() {
		err := http.ListenAndServe(":"+sPorta, api.Rotas)
		if err != nil {
			global.Logger.Println("ListenAndServe: ", err)
		}
	}()

	global.Logger.Println("Servidor Iniciado com sucesso Http: " + sPorta + " ...")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	global.Logger.Println("Finalizando servidor")
	os.Exit(0)

}
