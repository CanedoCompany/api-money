package handler

import (
	"fmt"
	"net/http"

	"github.com/CanedoCompany/api-money/internal/entity"
	"github.com/gin-gonic/gin"
)

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

func GetAllTransactionHandler(ctx *gin.Context) {
	transactions := []entity.Transaction{}

	if err := db.Find(&transactions).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listining transactions")
		return
	}

	sendSuccess(ctx, "list-transactions", transactions)
}
