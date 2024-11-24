package controllers

import (
	"digital-wallet-api/internal/models"
	"digital-wallet-api/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GroupWalletController defines the controller layer for group wallet operations
type GroupWalletController struct {
	GroupWalletService *service.GroupWalletService
}

// NewGroupWalletController creates a new GroupWalletController instance
func NewGroupWalletController(service *service.GroupWalletService) *GroupWalletController {
	return &GroupWalletController{GroupWalletService: service}
}

// CreateGroupWallet handles the creation of a new group wallet
func (gwc *GroupWalletController) CreateGroupWallet(ctx *gin.Context) {
	var groupWallet models.GroupWallet
	walletId, err := strconv.Atoi(ctx.Param("wallet_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	groupWallet.CreatedBy = walletId

	if err := ctx.ShouldBindJSON(&groupWallet); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := gwc.GroupWalletService.CreateGroupWallet(groupWallet); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Group wallet created successfully"})
}

// GetGroupWalletByID retrieves a group wallet by its ID
func (gwc *GroupWalletController) GetGroupWalletByID(ctx *gin.Context) {
	walletId, err := strconv.Atoi(ctx.Param("wallet_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wallet ID"})
		return
	}

	groupWallet, err := gwc.GroupWalletService.GetGroupWalletByID(walletId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if groupWallet.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Group wallet not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": groupWallet})
}

// UpdateGroupWalletGoal updates the goal of an existing group wallet
func (gwc *GroupWalletController) UpdateGroupWalletGoal(ctx *gin.Context) {
	walletId, err := strconv.Atoi(ctx.Param("wallet_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wallet ID"})
		return
	}

	var groupWallet models.GroupWallet
	if err := ctx.ShouldBindJSON(&groupWallet); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	groupWallet.ID = walletId
	if err := gwc.GroupWalletService.UpdateGroupWalletGoal(groupWallet); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Group wallet goal updated successfully"})
}

// DeleteGroupWallet handles the deletion of a group wallet
func (gwc *GroupWalletController) DeleteGroupWallet(ctx *gin.Context) {
	walletId, err := strconv.Atoi(ctx.Param("wallet_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wallet ID"})
		return
	}

	if err := gwc.GroupWalletService.DeleteGroupWallet(walletId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Group wallet deleted successfully"})
}
