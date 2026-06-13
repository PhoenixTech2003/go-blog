package api

import "net/http"

func TagsRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", ApiConfig.GetTags)
	mux.HandleFunc("POST /", ApiConfig.CreateTags)

	return mux
}
