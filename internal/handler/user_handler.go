package handler

import (
	"fmt"
	"net/http"

	"github.com/CanedoCompany/api-money/internal/entity"
	"github.com/gin-gonic/gin"
)

func CreateUserHandler(ctx *gin.Context) {
	request := CreateUserRequest{}

	ctx.BindJSON(&request)

	// Request error
	if err := request.Validate(); err != nil {
		logger.Errorf("validation err: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := entity.User{
		Name:       request.Name,
		Email:      request.Email,
		Password:   request.Password,
		Profission: request.Profission,
	}

	if err := db.Create(&user).Error; err != nil {
		logger.Errorf("error creating user: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating user on database")
		return
	}

	sendSuccess(ctx, "create-user", user)
}

func GetUserHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	user := entity.User{}

	if err := db.First(&user, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "user not found")
		return
	}

	sendSuccess(ctx, "show-user", user)
}

func EditUserHandler(ctx *gin.Context) {
	request := UpdateUserRequest{}

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

	user := entity.User{}

	if err := db.First(&user, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "user not found")
		return
	}

	// Update User

	if request.Name != "" {
		user.Name = request.Name
	}

	if request.Email != "" {
		user.Email = request.Email
	}

	if request.Password != "" {
		user.Password = request.Password
	}

	if request.Profission != "" {
		user.Profission = request.Profission
	}

	if err := db.Save(&user).Error; err != nil {
		logger.Errorf("error updating user: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating user")
		return
	}

	sendSuccess(ctx, "updating-user", user)
}

func DeleteUserHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	user := entity.User{}

	// Find User
	if err := db.First(&user, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("user with id: %s not found", id))
		return
	}
	// Delete User
	if err := db.Delete(&user).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting user with id: %s", id))
		return
	}

	sendSuccess(ctx, "delete-user", user)
}

func GetAllUserHandler(ctx *gin.Context) {
	users := []entity.User{}

	if err := db.Find(&users).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listining users")
		return
	}

	sendSuccess(ctx, "list-users", users)
}
