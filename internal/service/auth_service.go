package service

import (
	"database/sql"
	"digital-wallet-api/internal/models"
	"digital-wallet-api/internal/repository"
)

// GroupWalletService defines the service layer for group wallet operations
type AuthService struct {
	DB *sql.DB
}

// CreateUser creates a new user
func (as *AuthService) CreateUser(user models.UserAccount) error {
	return repository.CreateUser(as.DB, user)
}
