package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUserHandler(ctx *gin.Context) {
	request := CreateUserRequest{}

	ctx.BindJSON(&request)

	// Request error
	if err := request.Validate(); err != nil {
		logger.Errorf("validate err: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := db.Create(&request).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		return
	}
}

func GetUserHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET User",
	})
}

func EditUserHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "EDIT User",
	})
}

func DeleteUserHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "DELETE User",
	})
}

func GetAllUserHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Users",
	})
}
