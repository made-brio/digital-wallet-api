package controllers

import (
	"database/sql"
	"digital-wallet-api/internal/models"
	"digital-wallet-api/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// WalletController defines the controller layer for wallet operations
type WalletController struct {
	WalletService *service.WalletService
}

// NewWalletController creates a new WalletController instance
func NewWalletController(walletService *service.WalletService) *WalletController {
	return &WalletController{WalletService: walletService}
}

func (wc *WalletController) CreateWallet(ctx *gin.Context) {

	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := wc.WalletService.CreateWallet(userId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Wallet created successfully"})
}

// CheckBalance handles retrieving wallet balance
func (wc *WalletController) CheckBalance(ctx *gin.Context) {
	walletID, err := strconv.Atoi(ctx.Param("wallet_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wallet ID"})
		return
	}

	wallet, err := wc.WalletService.CheckBalance(walletID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Wallet not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": wallet})
}

func (wc *WalletController) GetInfoByWalletId(ctx *gin.Context) {
	walletID, err := strconv.Atoi(ctx.Param("wallet_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wallet ID"})
		return
	}

	wallet, err := wc.WalletService.GetInfoByWalletId(walletID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Wallet not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": wallet})
}

// TopUp handles adding funds to the wallet
func (wc *WalletController) TopUp(ctx *gin.Context) {
	walletID, err := strconv.Atoi(ctx.Param("wallet_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wallet ID"})
		return
	}
	var payload models.TopUpRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payload.WalletID = walletID

	wallet, err := wc.WalletService.TopUp(payload)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Wallet not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": wallet})
}

// FreezeWallet handles freezing the wallet
func (wc *WalletController) FreezeWallet(ctx *gin.Context) {
	walletID, err := strconv.Atoi(ctx.Param("wallet_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wallet ID"})
		return
	}

	err = wc.WalletService.FreezeWallet(walletID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Wallet frozen successfully"})
}

// UnfreezeWallet handles unfreezing the wallet
func (wc *WalletController) UnfreezeWallet(ctx *gin.Context) {
	walletID, err := strconv.Atoi(ctx.Param("wallet_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wallet ID"})
		return
	}

	err = wc.WalletService.UnfreezeWallet(walletID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Wallet unfrozen successfully"})
}
