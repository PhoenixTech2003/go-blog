package web

import (
	"net/http"
)

func IndexRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", WebConfig.HandleIndex)

	return mux
}
