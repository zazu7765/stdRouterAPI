package server

import "net/http"

type Route struct {
	Name    string
	Method  string
	Pattern string
	Handler http.HandlerFunc
}

func NewRouter(routes []Route) *http.ServeMux {
	router := http.NewServeMux()
	for _, route := range routes {
		router.HandleFunc(route.Pattern, route.Handler)
	}
	return router
}
