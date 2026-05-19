package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/PhoenixTech2003/go-blog/internal/utils"
	"github.com/charmbracelet/log"
)

type createTagRequestBody struct {
	Name string `json:"name"`
}

type tag struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type createTagSuccessResponseBody struct {
	Data    tag    `json:"data"`
	Message string `json:"message"`
}

// CreateTags godoc
//
// @Tags tags
// @Summary Creates a new tag
// @Description This endpoint is used to create a new tag
// @Accept json
// @Produce json
// @Param qoute body createTagRequestBody true "Tag Data"
// @Success 201 {object} createTagSuccessResponseBody "Tag Created Succefully"
// @Failure 500 {object} errResponseBody "Internal Server Error"
// @Router /v1/api/tags [post]
func (cfg *ApiCfg) CreateTags(w http.ResponseWriter, r *http.Request) {
	requestBody := createTagRequestBody{}
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&requestBody)
	if err != nil {
		log.Errorf("Failed to decode request body %s", err.Error())
		errRespBody := errResponseBody{
			Message: "Failed to add tag to the database",
		}
		dat, err := json.Marshal(errRespBody)
		if err != nil {
			log.Fatalf("Failed to marshal json %s", err.Error())
			return
		}

		utils.RespondWithJson(w, dat, http.StatusInternalServerError)
		return
	}

	tagName := requestBody.Name

	tagData, err := cfg.Db.CreateTag(r.Context(), tagName)
	if err != nil {
		log.Errorf("Failed to create tag with error %s", err.Error())
		errResp := errResponseBody{
			Message: "Failed to create tag",
		}
		dat, err := json.Marshal(errResp)
		if err != nil {
			log.Fatalf("failed to marshal json %s", err.Error())
			return
		}
		utils.RespondWithJson(w, dat, http.StatusInternalServerError)
	}

	tag := tag{
		Id:        tagData.ID.String(),
		Name:      tagData.Name,
		CreatedAt: nullTimeValue(tagData.CreatedAt),
		UpdatedAt: nullTimeValue(tagData.UpdatedAt),
	}

	tagDataResp := createTagSuccessResponseBody{
		Data: tag,
	}

	dat, err := json.Marshal(tagDataResp)
	if err != nil {
		log.Fatalf("Failed to marshal the tag data %s", err.Error())
		return
	}

	utils.RespondWithJson(w, dat, http.StatusInternalServerError)

}
