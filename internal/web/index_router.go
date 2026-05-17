package web

import (
	"net/http"
)

func IndexRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /fragments/random-quote", WebConfig.HandleRandomQuoteFragment)
	mux.HandleFunc("GET /", WebConfig.HandleIndex)

	return mux
}
