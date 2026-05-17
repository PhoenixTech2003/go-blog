package api

import (
	"crypto/rand"
	"database/sql"
	"encoding/json"
	"log"
	"log/slog"
	"math/big"
	"net/http"
	"time"

	"github.com/PhoenixTech2003/go-blog/internal/database"
	"github.com/PhoenixTech2003/go-blog/internal/utils"
)

type qouteRequestBody struct {
	Author    string `json:"author"`
	QuoteText string `json:"qoute_text"`
}

type errResponseBody struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
}

type qoute struct {
	Id        string    `json:"id"`
	Author    string    `json:"author"`
	QouteText string    `json:"qoute_text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type createQouteSuccessResponseBody struct {
	Data    qoute `json:"data"`
	Message any   `json:"message"`
}

func nullTimeValue(nt sql.NullTime) time.Time {
	if nt.Valid {
		return nt.Time
	}
	return time.Time{}
}

// CreateQoute godoc
//
// @Tags go-blog
// @Summary Creates a new Qoute
// @Description This endpoint is used to create new go-blog
// @Accept json
// @Produce json
// @Param qoute body qouteRequestBody true "Qoute Data"
// @Success 201 {object} createQouteSuccessResponseBody "Qoute Created"
// @Failure 500 {object} errResponseBody "Internal Server Error"
// @Router /v1/api/go-blog [post]
func (cfg *ApiCfg) CreateQoute(w http.ResponseWriter, r *http.Request) {
	requestBody := qouteRequestBody{}
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&requestBody)
	if err != nil {
		slog.Error("An error occured while decoding request body", "error", err)
		errorResponse := errResponseBody{
			Message: "An error occured while creating the qoute",
		}

		dat, err := json.Marshal(errorResponse)
		if err != nil {
			log.Fatal("Failed to marshal json when creating error response for the creating qoute")
			return
		}
		utils.RespondWithJson(w, dat, http.StatusInternalServerError)
		return

	}
	createQouteParams := database.CreateQouteParams{
		Author:    requestBody.Author,
		QouteText: requestBody.QuoteText,
	}

	qouteData, err := cfg.Db.CreateQoute(r.Context(), createQouteParams)
	if err != nil {
		slog.Error("An error occured while inserting qoute into the database", "error", err)
		errorResponse := errResponseBody{
			Message: "Failed to create qoute",
		}
		dat, err := json.Marshal(errorResponse)
		if err != nil {
			log.Fatal("Failed to marshal json during failed insert of qoute")
			return
		}
		utils.RespondWithJson(w, dat, http.StatusInternalServerError)
		return
	}

	successResponseBody := createQouteSuccessResponseBody{
		Data: qoute{
			Id:        qouteData.ID.String(),
			Author:    qouteData.Author,
			QouteText: qouteData.QouteText,
			CreatedAt: nullTimeValue(qouteData.CreatedAt),
			UpdatedAt: nullTimeValue(qouteData.UpdatedAt),
		},
	}

	dat, err := json.Marshal(successResponseBody)
	if err != nil {
		slog.Error("An error occured while marshalling a succeed response for creating a qoute", "error", err)
		errorResponse := errResponseBody{
			Message: "Failed to provide respnose after creating qoute",
		}
		dat, err := json.Marshal(errorResponse)
		if err != nil {
			log.Fatal("Failed to marshal json after successfull creating of qoute")
			return
		}
		utils.RespondWithJson(w, dat, http.StatusInternalServerError)
		return
	}

	utils.RespondWithJson(w, dat, http.StatusOK)

}

// GetRandomQoute godoc
//
// @Tags go-blog
// @Summary Gets a random qoute from the server
// @Description This endpoint is used to return random go-blog
// @Produce json
// @Success 200 {object} createQouteSuccessResponseBody "Qoute returned"
// @Failure 500 {object} errResponseBody "Internal Server Error"
// @Router /v1/api/go-blog/random [get]
func (cfg *ApiCfg) GetRandomQoute(w http.ResponseWriter, r *http.Request) {
	qouteData, err := cfg.Db.ListQuotes(r.Context())
	if err != nil {
		slog.Error("Failed to fetch qoute", "error", err.Error())
		errorResponse := errResponseBody{
			Message: "Failed to get random qoute",
		}

		dat, err := json.Marshal(errorResponse)
		if err != nil {
			log.Fatal("Failed to marshal json after fetching qoute")
			return

		}
		utils.RespondWithJson(w, dat, http.StatusInternalServerError)
		return
	}
	if len(qouteData) == 0 {
		errorResponse := errResponseBody{
			Message: "No quotes available",
		}
		dat, err := json.Marshal(errorResponse)
		if err != nil {
			log.Fatal("Failed to marshal empty quotes response")
			return
		}
		utils.RespondWithJson(w, dat, http.StatusNotFound)
		return
	}

	maxValue := len(qouteData)
	val, err := rand.Int(rand.Reader, big.NewInt(int64(maxValue)))
	if err != nil {
		log.Fatal("Failed to generate a random value")
		return
	}
	randomQoute := qouteData[val.Int64()]

	successResponseBody := createQouteSuccessResponseBody{
		Data: qoute{
			Id:        randomQoute.ID.String(),
			Author:    randomQoute.Author,
			QouteText: randomQoute.QouteText,
			CreatedAt: nullTimeValue(randomQoute.CreatedAt),
			UpdatedAt: nullTimeValue(randomQoute.UpdatedAt),
		},
	}

	dat, err := json.Marshal(successResponseBody)
	if err != nil {
		slog.Error("failed to marshal random qoute data response", "error", err.Error())
		return
	}
	utils.RespondWithJson(w, dat, http.StatusOK)
}
