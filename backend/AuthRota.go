package backend

import "net/http"

func (bc *backend) AuthRota(h http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// global.Logger.Printf("\n %+v\n", r.URL.)
		// fmt.Println("uri Host: " + r.URL.Host + " Scheme: " + r.RequestURI)
		// fmt.Println("Host: " + r.Host)

		// q := r.URL.Query()
		// for _, param := range params {
		// 	if len(q.Get(param)) == 0 {
		// 		http.Error(w, "missing "+param, http.StatusBadRequest)
		// 		return // exit early
		// 	}
		// }

		h.ServeHTTP(w, r) // all params present, proceed
	})
}
