package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	
)

func CreateUserHandler(ctx *gin.Context) {
  request := CreateUserRequest{}

  ctx.BindJSON(&request)

  if err := request.Validate(); err != nil {
	logger.Errorf("Validation error: %v", err.Error())
	sendError(ctx, http.StatusBadRequest, err.Error())
    return
  }
}