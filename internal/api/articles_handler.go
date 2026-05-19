package api

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/PhoenixTech2003/go-blog/internal/database"
	"github.com/PhoenixTech2003/go-blog/internal/utils"
)

type articleRequestBody struct {
	Author      string `json:"author"`
	ArticleText string `json:"article_text"`
	TagId       string `json:"tag_id"`
}

type article struct {
	Id          string    `json:"id"`
	Author      string    `json:"author"`
	ArticleText string    `json:"article_text"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	PublishedAt time.Time `json:"published_at"`
}

type createArticleSuccessResponseBody struct {
	Data    article `json:"data"`
	Message any     `json:"message"`
}

// CreateArticle godoc
//
// @Tags articles
// @Summary Creates a new blog article
// @Description This endpoint is used to create a new article
// @Accept json
// @Produce json
// @Param qoute body articleRequestBody true "Article Data"
// @Success 201 {object} createArticleSuccessResponseBody "Article Created Succefully"
// @Failure 500 {object} errResponseBody "Internal Server Error"
// @Router /v1/api/articles [post]
func (cfg *ApiCfg) CreateArticle(w http.ResponseWriter, r *http.Request) {
	requestBody := articleRequestBody{}
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&requestBody)
	if err != nil {
		slog.Error("An error occured while decoding request body", "error", err)
		errorResponse := errResponseBody{
			Message: "An error occured while creating the article",
		}

		dat, err := json.Marshal(errorResponse)
		if err != nil {
			log.Fatal("Failed to marshal json when creating error response for the creating article")
			return
		}
		utils.RespondWithJson(w, dat, http.StatusInternalServerError)
		return

	}
	createQouteParams := database.CreateAritcleParams{
		Author:      requestBody.Author,
		ArticleText: requestBody.ArticleText,
	}

	articleData, err := cfg.Db.CreateAritcle(r.Context(), createQouteParams)
	if err != nil {
		slog.Error("An error occured while inserting article into the database", "error", err)
		errorResponse := errResponseBody{
			Message: "Failed to create article",
		}
		dat, err := json.Marshal(errorResponse)
		if err != nil {
			log.Fatal("Failed to marshal json during failed insert of article")
			return
		}
		utils.RespondWithJson(w, dat, http.StatusInternalServerError)
		return
	}

	successResponseBody := createArticleSuccessResponseBody{
		Data: article{
			Id:          articleData.ID.String(),
			Author:      articleData.Author,
			ArticleText: articleData.ArticleText,
			CreatedAt:   nullTimeValue(articleData.CreatedAt),
			UpdatedAt:   nullTimeValue(articleData.UpdatedAt),
			PublishedAt: nullTimeValue(articleData.PublishedAt),
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
