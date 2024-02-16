package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	
)

func CreateUserHandler(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,gin.H{
			"msg": "POST User",
		})
}

func GetUserHandler(ctx *gin.Context) {
  ctx.JSON(http.StatusOK,gin.H{
    "msg": "GET User",
  })
}

func EditUserHandler(ctx *gin.Context) {
  ctx.JSON(http.StatusOK,gin.H{
    "msg": "EDIT User",
  })
}

func DeleteUserHandler(ctx *gin.Context) {
  ctx.JSON(http.StatusOK,gin.H{
    "msg": "DELETE User",
  })
}

func GetAllUserHandler(ctx *gin.Context) {
  ctx.JSON(http.StatusOK,gin.H{
    "msg": "GET Users",
  })
}