package api

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/PhoenixTech2003/go-blog/internal/database"
	"github.com/PhoenixTech2003/go-blog/internal/utils"
	"github.com/google/uuid"
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

type getArticlesSuccessResponseBody struct {
	Data    []article `json:"data"`
	Message string    `json:"message"`
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

	parsedUUID := uuid.MustParse(requestBody.TagId)
	createTagArticleParams := database.AddTagToArticleParams{
		AritcleID: articleData.ID,
		TagID:     parsedUUID,
	}
	_, err = cfg.Db.AddTagToArticle(r.Context(), createTagArticleParams)
	if err != nil {
		slog.Error("An error occured while adding tag to article into the database", "error", err)
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

// GetArticles godoc
//
// @Tags articles
// @Summary Gets all articles
// @Description This endpoint is used to get all articles
// @Produce json
// @Success 201 {object} getArticlesSuccessResponseBody "Articles fetched  Succefully"
// @Failure 500 {object} errResponseBody "Internal Server Error"
// @Router /v1/api/articles [get]
func (cfg *ApiCfg) GetArticles(w http.ResponseWriter, r *http.Request) {
	dat, err := cfg.Db.ListArticles(r.Context())
	if err != nil {
		errRespnse := errResponseBody{
			Message: "Failed to fetch all articles",
		}
		dat, err := json.Marshal(errRespnse)
		if err != nil {
			log.Fatal("Failed to marshal json")
			return
		}
		utils.RespondWithJson(w, dat, http.StatusInternalServerError)
	}

	articlesList := make([]article, 0)

	for _, articleDetails := range dat {
		articleData := article{
			Id:          articleDetails.ID.String(),
			Author:      articleDetails.Author,
			ArticleText: articleDetails.ArticleText,
			CreatedAt:   nullTimeValue(articleDetails.CreatedAt),
			PublishedAt: nullTimeValue(articleDetails.PublishedAt),
			UpdatedAt:   nullTimeValue(articleDetails.UpdatedAt),
		}

		articlesList = append(articlesList, articleData)
	}

	successBody := getArticlesSuccessResponseBody{
		Data: articlesList,
	}
	successBodyDat, err := json.Marshal(successBody)
	if err != nil {
		errRespnse := errResponseBody{
			Message: "Failed to fetch all articles",
		}
		dat, err := json.Marshal(errRespnse)
		if err != nil {
			log.Fatal("Failed to marshal json")
			return
		}
		utils.RespondWithJson(w, dat, http.StatusInternalServerError)
	}

	utils.RespondWithJson(w, successBodyDat, http.StatusOK)

}
