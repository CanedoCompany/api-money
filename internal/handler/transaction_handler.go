package handler

import (
	"fmt"
	"net/http"

	"github.com/CanedoCompany/api-money/internal/entity"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Sumary Create Transaction
// @Description Create a new Transaction
// @Tags Transactions
// @Accept json
// @Produce json
// @Param request body CreateTransactionRequest true "Request body"
// @Success 200 {object} CreateTransactionResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /transaction [post]
func CreateTransactionHandler(ctx *gin.Context) {
	request := CreateTransactionRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation err: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	transaction := entity.Transaction{
		NameTransaction: request.NameTransaction,
		Value:           request.Value,
		Category:        request.Category,
	}

	if err := db.Create(&transaction).Error; err != nil {
		logger.Errorf("error creating transaction: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating transaction on database")
		return
	}

	sendSuccess(ctx, "create-transaction", transaction)
}

// @BasePath /api/v1

// @Sumary Show Transaction
// @Description Show a new Transaction
// @Tags Transactions
// @Accept json
// @Produce json
// @Param id query string true "Transaction identification"
// @Success 200 {object} ShowTransactionResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /transaction [get]
func GetTransactionHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	transaction := entity.Transaction{}

	if err := db.First(&transaction, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "Transaction not found")
		return
	}

	sendSuccess(ctx, "show-transaction", transaction)
}

// @BasePath /api/v1

// @Sumary Update Transaction
// @Description Update an Transaction
// @Tags Transactions
// @Accept json
// @Produce json
// @Param id query string true "Transaction identification"
// @Param request body UpdateTransactionRequest true "Transaction data to Update"
// @Success 200 {object} CreateTransactionResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /transaction [put]
func EditTransactionHandler(ctx *gin.Context) {
	request := UpdateTransactionRequest{}

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

	transaction := entity.Transaction{}

	if err := db.First(&transaction, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "Transaction not found")
		return
	}

	if request.NameTransaction != "" {
		transaction.NameTransaction = request.NameTransaction
	}

	if request.Value <= 0 {
		transaction.Value = request.Value
	}

	if request.Category != "" {
		transaction.Category = request.Category
	}

	if err := db.Save(&transaction).Error; err != nil {
		logger.Errorf("error updating transaction: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating transaction")
		return
	}

	sendSuccess(ctx, "updating-transaction", transaction)
}

// @BasePath /api/v1

// @Sumary Delete Transaction
// @Description Delete a new Transaction
// @Tags Transactions
// @Accept json
// @Produce json
// @Param id query string true "Transaction identification"
// @Success 200 {object} DeleteTransactionResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /transaction [delete]
func DeleteTransactionHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	transaction := entity.Transaction{}

	if err := db.First(&transaction, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("user with id: %s not found", id))
		return
	}

	if err := db.Delete(&transaction).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting transaction with id: %s", id))
		return
	}

	sendSuccess(ctx, "delete-transaction", transaction)
}

// @BasePath /api/v1

// @Sumary Get All Transactions
// @Description List all Transactions
// @Tags Transactions
// @Accept json
// @Produce json
// @Success 200 {object} GetAllTransactionResponse
// @Failure 500 {object} ErrorResponse
// @Router /transactions [get]
func GetAllTransactionHandler(ctx *gin.Context) {
	transactions := []entity.Transaction{}

	if err := db.Find(&transactions).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listining transactions")
		return
	}

	sendSuccess(ctx, "list-transactions", transactions)
}
