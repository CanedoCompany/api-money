package router

import (
	// "net/http"
	"github.com/CanedoCompany/api-money/internal/handler"
	"github.com/gin-gonic/gin"
)

// Routes
// "/api/v1" here we don't needa version api
func initializeRoutes(router *gin.Engine) {

	handler.InitializeHandler()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/user", handler.GetUserHandler)

		v1.POST("/user", handler.CreateUserHandler)

		v1.PUT("/user", handler.EditUserHandler)

		v1.DELETE("/user", handler.DeleteUserHandler)

		v1.GET("/users", handler.GetAllUserHandler)

		v1.GET("/transaction", handler.GetTransactionHandler)

		v1.POST("/transaction", handler.CreateTransactionHandler)

		v1.PUT("/transaction", handler.EditTransactionHandler)

		v1.DELETE("/transaction", handler.DeleteTransactionHandler)

		v1.GET("/transactions", handler.GetAllTransactionHandler)
	}

}
