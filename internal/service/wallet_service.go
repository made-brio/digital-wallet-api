package service

import (
	"database/sql"
	"digital-wallet-api/internal/models"
	"digital-wallet-api/internal/repository"
)

// WalletService defines the service layer for wallet operations
type WalletService struct {
	DB *sql.DB
}

// NewWalletService creates a new WalletService instance
func NewWalletService(db *sql.DB) *WalletService {
	return &WalletService{DB: db}
}

// CreateW Wallet Balance
func (s *WalletService) CreateWallet(userId int) error {
	return repository.CreateWallet(s.DB, userId)
}

// CheckBalance retrieves the wallet balance
func (s *WalletService) CheckBalance(walletID int) (models.Wallet, error) {
	return repository.CheckBalance(s.DB, walletID)
}

func (s *WalletService) GetInfoByWalletId(walletID int) (models.Wallet, error) {
	return repository.GetInfoByWalletId(s.DB, walletID)
}

// TopUp adds funds to the wallet
func (s *WalletService) TopUp(payload models.TopUpRequest) (models.Wallet, error) {
	return repository.TopUp(s.DB, payload)
}

// FreezeWallet freezes the wallet
func (s *WalletService) FreezeWallet(walletID int) error {
	return repository.FreezeWallet(s.DB, walletID)
}

// UnfreezeWallet unfreezes the wallet
func (s *WalletService) UnfreezeWallet(walletID int) error {
	return repository.UnfreezeWallet(s.DB, walletID)
}
