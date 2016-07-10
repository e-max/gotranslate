package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"log"
)

var routes = []struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}{
	{"Index", "GET", "/", GetIndex},
	{"AddWord", "POST", "/add", AddWordPost},
	{"TranslateWord", "GET", "/translate/{olocale:[a-z]{2}}/{title:[A-Za-z]+}/{tlocale:[a-z]{2}}", TranslateWord},
}

// Decorate http handlers with logging
func Logger(handler http.Handler, name string) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			handler.ServeHTTP(w, r)
			log.Printf("%s\t%s\t%s\t%s", r.Method, r.RequestURI, name, time.Since(start))
	})
}

// Create router for Api
func CreateRouter() *mux.Router {
	r := mux.NewRouter()

	for _, route := range routes {
		handler := Logger(route.HandlerFunc, route.Name)
		r.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)
	}

	return r
}
