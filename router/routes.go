package router

import (
	// "net/http"
	docs "github.com/CanedoCompany/api-money/docs"
	"github.com/CanedoCompany/api-money/internal/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Routes
// "/api/v1" here we don't needa version api
func initializeRoutes(router *gin.Engine) {

	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
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

		v1.GET("/cards", handler.GetAllCardHandler)

		v1.GET("/card", handler.GetCardHandler)

		v1.POST("/card", handler.CreateCardHandler)

		v1.PUT("/card", handler.EditCardHandler)

		v1.DELETE("/card", handler.DeleteCardHandler)

		v1.GET("/banks", handler.GetAllBankHandler)

		v1.GET("/bank", handler.GetBankHandler)

		v1.POST("/bank", handler.CreateBankHandler)

		v1.PUT("/bank", handler.EditBankhandler)

		v1.DELETE("/bank", handler.DeleteBankHandler)
	}

	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}
