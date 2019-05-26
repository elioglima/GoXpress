package backend

import (
	"fmt"
	"goxpress/backend/assistente"
	"goxpress/backend/global"
	ApiRT "goxpress/backend/rotas/api"
	"goxpress/backend/rotas/publico"
	pub_usuario "goxpress/backend/rotas/publico/usuario"
	"net/http"
	"strings"

	"goxpress/backend/database/rotas_http"

	"github.com/gorilla/mux"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	UrlHttp   string
	Todos     []Todo
}

func (bc *backend) Geral() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		global.Logger.Sucesso("\n %+v\n", r.URL)
	})
}

func NewRouter() *mux.Router {
	dirPublic := global.DirPublic()
	router := mux.NewRouter().StrictSlash(true)

	// Server CSS, JS & Images Statically.
	router.PathPrefix(dirPublic).
		Handler(http.StripPrefix(dirPublic, http.FileServer(http.Dir("."+dirPublic))))

	return router
}

func (bc *backend) LogRotas() {
	err := bc.Rotas.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			fmt.Println("Rota:", pathTemplate)
		}
		pathRegexp, err := route.GetPathRegexp()
		if err == nil {
			fmt.Println("Endereço regexp:", pathRegexp)
		}
		queriesTemplates, err := route.GetQueriesTemplates()
		if err == nil {
			fmt.Println("Consulta templates:", strings.Join(queriesTemplates, ","))
		}
		queriesRegexps, err := route.GetQueriesRegexp()
		if err == nil {
			fmt.Println("Consulta regexps:", strings.Join(queriesRegexps, ","))
		}
		methods, err := route.GetMethods()
		if err == nil {
			fmt.Println("Meodo:", strings.Join(methods, ","))
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}
}

func (bc *backend) setRotas() {

	bc.Rotas = NewRouter()
	bc.DataView = NewDataView()
	bc.LogRotas()
	setAuthMiddleware(bc.Rotas)

	bc.Rotas.NotFoundHandler = http.HandlerFunc(publico.PageNotFound)
	bc.Rotas.HandleFunc("/", publico.Home).Methods("GET")
	bc.Rotas.HandleFunc("/usuario/acesso", pub_usuario.Acesso).Methods("GET")
	bc.Rotas.HandleFunc("/usuario/acesso", pub_usuario.Logar).Methods("POST")
	bc.Rotas.HandleFunc("/api/acesso", ApiRT.Acesso).Methods("POST")
	bc.Rotas.HandleFunc("/api/info", ApiRT.Info).Methods("POST")

	bc.Rotas.HandleFunc("/db/empresa", rotas_http.RT_empresa).Methods("GET")
	bc.Rotas.HandleFunc("/db/assistente", rotas_http.RT_assistente).Methods("GET")
	bc.Rotas.HandleFunc("/db/assistente/apresentacao", rotas_http.RT_apresentacao).Methods("GET")
	bc.Rotas.HandleFunc("/db/assistente/resposta_acao/add", rotas_http.RT_resposta_acao_add).Methods("POST")
	bc.Rotas.HandleFunc("/db/assistente/resposta_acao/edd", rotas_http.RT_resposta_acao_edd).Methods("POST")

	bc.Rotas.HandleFunc("/assistente", assistente.RT_assistente).Methods("POST")

}
