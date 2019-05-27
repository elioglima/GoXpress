package src

import (
	"GoXpress/api/src/global"
	"fmt"
	"net/http"
	"strings"

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

func (api *api) Geral() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		global.Logger.Println("\n %+v\n", r.URL)
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

func (api *api) LogRotas() {
	err := api.Rotas.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
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

func (api *api) setRotas() {

	api.Rotas = NewRouter()
	api.LogRotas()
	// setAuthMiddleware(api.Rotas)

	api.Rotas.HandleFunc("/", routes)
	api.Rotas.HandleFunc("/{nivel1}", routes)
	api.Rotas.HandleFunc("/{nivel1}/{nivel2}", routes)
	api.Rotas.HandleFunc("/{nivel1}/{nivel2}/{nivel3}", routes)
	api.Rotas.HandleFunc("/{nivel1}/{nivel2}/{nivel3}/{nivel4}", routes)
	api.Rotas.HandleFunc("/{nivel1}/{nivel2}/{nivel3}/{nivel4}/{nivel5}", routes)
}

// func GetStructByName(name string) interface{} {
// 	return reflect.ValueOf(typeRegistry[name]).Interface()
// }

func routes(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	nivel1 := params["nivel1"]
	nivel2 := params["nivel2"]
	nivel3 := params["nivel3"]
	nivel4 := params["nivel4"]
	nivel5 := params["nivel5"]

	fmt.Println("rotas", nivel1, nivel2, nivel3, nivel4, nivel5)

}
