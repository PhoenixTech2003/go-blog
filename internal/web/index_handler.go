package web

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"log"
	"log/slog"
	"math/big"
	"net/http"

	"github.com/starfederation/datastar-go/datastar"

	"github.com/PhoenixTech2003/go-blog/internal/database"
	"github.com/PhoenixTech2003/go-blog/internal/utils"
	"github.com/PhoenixTech2003/go-blog/internal/web/ui/components"
	"github.com/PhoenixTech2003/go-blog/internal/web/ui/pages"
)

type errResponseBody struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func (cfg *WebCfg) pickRandomQuote(ctx context.Context) (database.Quote, error) {
	qoutesData, err := cfg.Db.ListQuotes(ctx)
	if err != nil {
		return database.Quote{}, err
	}
	if len(qoutesData) == 0 {
		return database.Quote{}, nil
	}
	val, err := rand.Int(rand.Reader, big.NewInt(int64(len(qoutesData))))
	if err != nil {
		return database.Quote{}, err
	}
	return qoutesData[val.Int64()], nil
}

func (cfg *WebCfg) HandleIndex(w http.ResponseWriter, r *http.Request) {
	randomQoute, err := cfg.pickRandomQuote(r.Context())
	if err != nil {
		slog.Error("Failed to fetch qoutes", "error", err.Error())
		errorResponse := errResponseBody{
			Message: "Failed to get random qoute",
		}

		dat, err := json.Marshal(errorResponse)
		if err != nil {
			log.Fatal("Failed to marshal json after fetching qoutes")
			return

		}
		utils.RespondWithJson(w, dat, http.StatusInternalServerError)
		return
	}

	component := pages.Index(randomQoute)
	err = component.Render(r.Context(), w)
	if err != nil {
		slog.Error("failed to load index page", "error", err.Error())
	}
}

// HandleRandomQuoteFragment streams a Datastar SSE patch for the #random island.
func (cfg *WebCfg) HandleRandomQuoteFragment(w http.ResponseWriter, r *http.Request) {
	randomQoute, err := cfg.pickRandomQuote(r.Context())
	if err != nil {
		slog.Error("Failed to fetch qoutes for fragment", "error", err.Error())
		http.Error(w, "Failed to fetch quotes", http.StatusInternalServerError)
		return
	}

	sse := datastar.NewSSE(w, r)
	frag := components.RandomQouteSection(randomQoute)
	if err := sse.PatchElementTempl(frag); err != nil {
		slog.Error("failed to stream random quote fragment", "error", err.Error())
	}
}
