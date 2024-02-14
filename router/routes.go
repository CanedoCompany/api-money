package router

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


// Routes
// "/api/v1" here we don't needa version api
func initializeRoutes(router *gin.Engine) {

 v1 := router.Group("/api/v1")
 {
	v1.GET("/user", func(ctx *gin.Context){
		ctx.JSON(http.StatusOK,gin.H{
			"msg": "GET User",
		})
	})

	v1.POST("/user", func(ctx *gin.Context){
		ctx.JSON(http.StatusOK,gin.H{
			"msg": "CREATE User",
		})
	})

	v1.PUT("/user", func(ctx *gin.Context){
		ctx.JSON(http.StatusOK,gin.H{
			"msg": "EDIT User",
		})
	})

	v1.DELETE("/user", func(ctx *gin.Context){
		ctx.JSON(http.StatusOK,gin.H{
			"msg": "DELETE User",
		})
	})

	v1.GET("/users", func(ctx *gin.Context){
		ctx.JSON(http.StatusOK,gin.H{
			"msg": "GET All Users",
		})
	})
 }

}