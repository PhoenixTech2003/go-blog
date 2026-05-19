package api

import "net/http"

func ArticlesRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /", ApiConfig.CreateArticle)

	return mux
}
