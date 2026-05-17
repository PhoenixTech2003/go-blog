package api

import "net/http"

func QouteRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /", ApiConfig.CreateQoute)
	mux.HandleFunc("GET /random", ApiConfig.GetRandomQoute)

	return mux
}
