package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

// Route structure
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		GetIndex,
	},
	Route{
		"AddWord",
		"POST",
		"/add",
		AddWordPost,
	},
	Route{
		"AddWord",
		"GET",
		"/add/{olocale}/{oword}/{tlocale}/{tword}",
		AddWordGet,
	},
	Route{
		"TranslateWord",
		"GET",
		"/translate/{locale}/{word}",
		TranslateWord,
	},
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
