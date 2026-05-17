package web

import "net/http"

func ComposeRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", WebConfig.HandleCompose)
	mux.HandleFunc("POST /qoutes/", WebConfig.HandleCreateQoute)

	return mux
}
