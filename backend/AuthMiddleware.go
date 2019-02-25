package backend

import (
	"goxpress/backend/global"
	"net/http"

	"github.com/gorilla/mux"
)

// Define our struct
type AuthMiddleware struct {
	TokenUsers map[string]string
}

// Initialize it somewhere
func (amw *AuthMiddleware) Populate() {
	// amw.TokenUsers["00000000"] = "user0"
	// amw.TokenUsers["aaaaaaaa"] = "userA"
	// amw.TokenUsers["05f717e5"] = "randomUser"
	// amw.TokenUsers["deadbeef"] = "user0"
}

// Middleware function, which will be called for each request
func (amw *AuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// r.Host

		global.SetHttpHome("http://" + r.Host)
		global.SetHttpSolicitado("http://" + r.Host + r.URL.RequestURI())
		// global.Logger.Rosa(global.HttpSolicitado())

		//fmt.Printf("%C \n", r.Header)
		next.ServeHTTP(w, r)
		// if user, found := amw.TokenUsers[token]; found {
		// 	// We found the token in our map
		// 	log.Printf("Authenticated user %s\n", user)
		// 	// Pass down the request to the next middleware (or final handler)
		// 	next.ServeHTTP(w, r)
		// } else {
		// 	// Write an error and stop the handler chain
		// 	http.Error(w, "Forbidden", http.StatusForbidden)
		// }
	})
}

func setAuthMiddleware(Rotas *mux.Router) {
	amw := AuthMiddleware{}
	amw.Populate()
	Rotas.Use(amw.Middleware)
}
