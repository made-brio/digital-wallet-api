package service

import (
	"database/sql"
	"digital-wallet-api/internal/models"
	"digital-wallet-api/internal/repository"
)

// TransactionService defines the service layer for transaction operations
type TransactionService struct {
	DB *sql.DB
}

// NewTransactionService creates a new TransactionService instance
func NewTransactionService(db *sql.DB) *TransactionService {
	return &TransactionService{DB: db}
}

// Transfer handles fund transfers between wallets
func (s *TransactionService) Transfer(transfer models.TransferRequest) error {
	return repository.Transfer(s.DB, transfer)
}

// IncomeHistory retrieves all income transactions for a wallet
func (s *TransactionService) IncomeHistory(walletID int) ([]models.Transaction, error) {
	return repository.IncomeHistory(s.DB, walletID)
}

// ExpenseHistory retrieves all expense transactions for a wallet
func (s *TransactionService) ExpenseHistory(walletID int) ([]models.Transaction, error) {
	return repository.ExpenseHistory(s.DB, walletID)
}

// TransactionsHistory retrieves all transactions for a wallet
func (s *TransactionService) TransactionsHistory(walletID int) ([]models.Transaction, error) {
	return repository.TransactionsHistory(s.DB, walletID)
}
