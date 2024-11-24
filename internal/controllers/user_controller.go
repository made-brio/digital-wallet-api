package controllers

import (
	"digital-wallet-api/internal/models"
	"digital-wallet-api/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserController defines the controller layer for user operations
type UserController struct {
	UserService *service.UserService
}

// NewUserController creates a new UserController instance
func NewUserController(userService *service.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (uc *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := uc.UserService.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"result": users})
}

func (uc *UserController) GetUserByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := uc.UserService.GetUserByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": user})
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var user models.UserAccount

	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user.ID = userId
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uc.UserService.UpdateUser(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
func (uc *UserController) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user := models.UserAccount{ID: id}
	if err := uc.UserService.DeleteUser(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})

}
