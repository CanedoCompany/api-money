package handler

import (
	"fmt"
	"net/http"

	"github.com/CanedoCompany/api-money/internal/entity"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Sumary Create Card
// @Description Create a new Card
// @Tags Cards
// @Accept json
// @Produce json
// @Param request body CreateCardRequest true "Request body"
// @Success 200 {object} CreateCardResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /cards [post]
func CreateCardHandler(ctx *gin.Context) {
	request := CreateCardRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation err: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	card := entity.Card{
		NameCard:   request.NameCard,
		NameOwner:  request.NameOwner,
		NumberCard: request.NumberCard,
		CvvCard:    request.CvvCard,
	}

	if err := db.Create(&card).Error; err != nil {
		logger.Errorf("error creating card: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating card on database")
		return
	}

	sendSuccess(ctx, "create-card", card)
}

// @BasePath /api/v1

// @Sumary Show Card
// @Description Show a new Card
// @Tags Cards
// @Accept json
// @Produce json
// @Param id query string true "Card identification"
// @Success 200 {object} ShowCardResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /card [get]
func GetCardHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	card := entity.Card{}

	if err := db.First(&card, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "card not found")
		return
	}

	sendSuccess(ctx, "show-card", card)
}

// @BasePath /api/v1

// @Sumary Update Card
// @Description Update an Card
// @Tags Cards
// @Accept json
// @Produce json
// @Param id query string true "Card identification"
// @Param request body UpdateCardRequest true "Card data to Update"
// @Success 200 {object} CreateCardResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /card [put]
func EditCardHandler(ctx *gin.Context) {
	request := UpdateCardRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	card := entity.Card{}

	if err := db.First(&card, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "card not found")
		return
	}

	if request.NameCard != "" {
		card.NameCard = request.NameCard
	}

	if request.NameOwner != "" {
		card.NameOwner = request.NameOwner
	}

	if request.NumberCard <= 0 {
		card.NumberCard = request.NumberCard
	}

	if request.CvvCard <= 0 {
		card.CvvCard = request.CvvCard
	}

	if err := db.Save(&card).Error; err != nil {
		logger.Errorf("error updating card: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating card")
		return
	}

	sendSuccess(ctx, "updating-card", card)
}

// @BasePath /api/v1

// @Sumary Delete Card
// @Description Delete a new Card
// @Tags Cards
// @Accept json
// @Produce json
// @Param id query string true "Card identification"
// @Success 200 {object} DeleteCardResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /card [delete]
func DeleteCardHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	card := entity.Card{}

	if err := db.First(&card, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("card with id: %s not found", id))
		return
	}

	if err := db.Delete(&card).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting card with id: %s", id))
		return
	}

	sendSuccess(ctx, "delete-card", card)
}

func GetAllCardHandler(ctx *gin.Context) {
	cards := []entity.Card{}

	if err := db.Find(&cards).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listining cards")
		return
	}

	sendSuccess(ctx, "list-cards", cards)
}
