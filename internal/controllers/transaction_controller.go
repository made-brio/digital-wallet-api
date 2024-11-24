package controllers

import (
	"database/sql"
	"digital-wallet-api/internal/models"
	"digital-wallet-api/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TransactionController defines the controller layer for transaction operations
type TransactionController struct {
	TransactionService *service.TransactionService
}

// NewTransactionController creates a new TransactionController instance
func NewTransactionController(transactionService *service.TransactionService) *TransactionController {
	return &TransactionController{TransactionService: transactionService}
}

// Transfer handles fund transfers between wallets
func (c *TransactionController) Transfer(ctx *gin.Context) {
	var transfer models.TransferRequest
	if err := ctx.ShouldBindJSON(&transfer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.TransactionService.Transfer(transfer)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Insufficient funds or wallet not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Transfer completed successfully"})
}

// IncomeHistory retrieves all income transactions for a wallet
func (c *TransactionController) IncomeHistory(ctx *gin.Context) {
	walletID, err := strconv.Atoi(ctx.Param("wallet_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wallet ID"})
		return
	}

	transactions, err := c.TransactionService.IncomeHistory(walletID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": transactions})
}

// ExpenseHistory retrieves all expense transactions for a wallet
func (c *TransactionController) ExpenseHistory(ctx *gin.Context) {
	walletID, err := strconv.Atoi(ctx.Param("wallet_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wallet ID"})
		return
	}

	transactions, err := c.TransactionService.ExpenseHistory(walletID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": transactions})
}

// TransactionsHistory retrieves all transactions for a wallet
func (c *TransactionController) TransactionsHistory(ctx *gin.Context) {
	walletID, err := strconv.Atoi(ctx.Param("wallet_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wallet ID"})
		return
	}

	transactions, err := c.TransactionService.TransactionsHistory(walletID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": transactions})
}
