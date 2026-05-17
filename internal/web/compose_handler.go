package web

import (
	"log/slog"
	"net/http"

	"github.com/starfederation/datastar-go/datastar"

	"github.com/PhoenixTech2003/go-blog/internal/database"
	"github.com/PhoenixTech2003/go-blog/internal/web/ui/components"
	"github.com/PhoenixTech2003/go-blog/internal/web/ui/pages"
)

func (cfg *WebCfg) HandleCompose(w http.ResponseWriter, r *http.Request) {
	component := pages.Compose(nil)
	err := component.Render(r.Context(), w)
	if err != nil {
		slog.Error("failed to load compose page", "error", err.Error())
	}
}

func (cfg *WebCfg) HandleCreateQoute(w http.ResponseWriter, r *http.Request) {
	// ParseForm must run before NewSSE: flushing the SSE response can close the request body,
	// which breaks ParseForm (see datastar.ReadSignals docs: SSE after reading the body).
	if err := r.ParseForm(); err != nil {
		slog.Error("Failed to parse form value", "error", err.Error())
		sse := datastar.NewSSE(w, r)
		form := components.ComposeForm(new(false))
		if err := sse.PatchElementTempl(form); err != nil {
			slog.Error("Failed to stream compose form", "error", err.Error())
		}
		return
	}

	author := r.FormValue("author")
	qoute := r.FormValue("qoute")
	createQouteParams := database.CreateQouteParams{
		QouteText: qoute,
		Author:    author,
	}

	_, err := cfg.Db.CreateQoute(r.Context(), createQouteParams)
	if err != nil {
		slog.Error("Failed to create qoute with error %s", "error", err.Error())
		sse := datastar.NewSSE(w, r)
		form := components.ComposeForm(new(false))
		if err := sse.PatchElementTempl(form); err != nil {
			slog.Error("Failed to stream compose form", "error", err.Error())
		}
		return
	}

	sse := datastar.NewSSE(w, r)
	form := components.ComposeForm(new(true))
	if err := sse.PatchElementTempl(form); err != nil {
		slog.Error("Failed to stream compose form", "error", err.Error())
	}
}
