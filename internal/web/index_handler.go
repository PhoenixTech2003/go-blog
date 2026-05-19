package web

import (
	"log/slog"
	"net/http"

	"github.com/PhoenixTech2003/go-blog/internal/web/ui/pages"
)

type errResponseBody struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func (cfg *WebCfg) HandleIndex(w http.ResponseWriter, r *http.Request) {

	component := pages.Index("hello world")
	err := component.Render(r.Context(), w)
	if err != nil {
		slog.Error("failed to load index page", "error", err.Error())
	}
}
