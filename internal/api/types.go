package api

type errResponseBody struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
}
