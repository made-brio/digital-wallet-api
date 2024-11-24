package controllers

import (
	"digital-wallet-api/internal/models"
	"digital-wallet-api/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GroupWalletMemberController defines the controller layer for group wallet member operations
type GroupWalletMemberController struct {
	GroupWalletMemberService *service.GroupWalletMemberService
}

// NewGroupWalletMemberController creates a new GroupWalletMemberController instance
func NewGroupWalletMemberController(service *service.GroupWalletMemberService) *GroupWalletMemberController {
	return &GroupWalletMemberController{GroupWalletMemberService: service}
}

// AddMember adds a new member to a group wallet
func (gwmc *GroupWalletMemberController) AddMember(ctx *gin.Context) {
	walletID, err := strconv.Atoi(ctx.Param("wallet_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wallet ID"})
		return
	}

	var member models.GroupWalletMember
	if err := ctx.ShouldBindJSON(&member); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	member.GroupWalletID = walletID
	if err := gwmc.GroupWalletMemberService.AddMember(member); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Member added successfully"})
}

// GetMembers retrieves all members of a group wallet
func (gwmc *GroupWalletMemberController) GetAllMembers(ctx *gin.Context) {
	walletID, err := strconv.Atoi(ctx.Param("wallet_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wallet ID"})
		return
	}

	members, err := gwmc.GroupWalletMemberService.GetAllMembers(walletID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": members})
}

// UpdateMemberContribution updates a member's contribution to the group wallet
func (gwmc *GroupWalletMemberController) UpdateMemberContribution(ctx *gin.Context) {
	var member models.GroupWalletMember
	walletID, err := strconv.Atoi(ctx.Param("wallet_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wallet ID"})
		return
	}

	memberID, err := strconv.Atoi(ctx.Param("member_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid member ID"})
		return
	}

	member.GroupWalletID = walletID
	member.UserID = memberID
	if err := ctx.ShouldBindJSON(&member); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := gwmc.GroupWalletMemberService.UpdateMemberContribution(member); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Member contribution updated successfully"})
}

// RemoveMember removes a member from the group wallet
func (gwmc *GroupWalletMemberController) RemoveMember(ctx *gin.Context) {
	walletID, err := strconv.Atoi(ctx.Param("wallet_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wallet ID"})
		return
	}

	memberID, err := strconv.Atoi(ctx.Param("member_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid member ID"})
		return
	}

	if err := gwmc.GroupWalletMemberService.RemoveMember(walletID, memberID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Member deleted successfully"})
}
