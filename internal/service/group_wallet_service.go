package service

import (
	"database/sql"
	"digital-wallet-api/internal/models"
	"digital-wallet-api/internal/repository"
)

// GroupWalletService defines the service layer for group wallet operations
type GroupWalletService struct {
	DB *sql.DB
}

// NewGroupWalletService creates a new instance of GroupWalletService
func NewGroupWalletService(db *sql.DB) *GroupWalletService {
	return &GroupWalletService{DB: db}
}

// CreateGroupWallet creates a new group wallet
func (gws *GroupWalletService) CreateGroupWallet(groupWallet models.GroupWallet) error {
	return repository.CreateGroupWallet(gws.DB, groupWallet)
}

// GetGroupWalletByID retrieves a group wallet by its ID
func (gws *GroupWalletService) GetGroupWalletByID(walletID int) (models.GroupWallet, error) {
	return repository.GetGroupWalletByID(gws.DB, walletID)
}

// UpdateGroupWalletGoal updates the goal of an existing group wallet
func (gws *GroupWalletService) UpdateGroupWalletGoal(groupWallet models.GroupWallet) error {
	return repository.UpdateGroupWalletGoal(gws.DB, groupWallet)
}

// DeleteGroupWallet deletes a group wallet by its ID
func (gws *GroupWalletService) DeleteGroupWallet(walletID int) error {
	return repository.DeleteGroupWallet(gws.DB, walletID)
}
