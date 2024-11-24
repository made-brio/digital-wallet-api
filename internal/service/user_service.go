package service

import (
	"database/sql"
	"digital-wallet-api/internal/models"
	"digital-wallet-api/internal/repository"
)

// UserService defines the service layer for user operations
type UserService struct {
	DB *sql.DB
}

// NewUserService creates a new UserService instance
func NewUserService(db *sql.DB) *UserService {
	return &UserService{DB: db}
}

// GetAllUsers retrieves all users
func (s *UserService) GetAllUsers() ([]models.UserAccount, error) {
	return repository.GetAllUsers(s.DB)
}

// GetUserByID retrieves a user by ID
func (s *UserService) GetUserByID(id int) (models.UserAccount, error) {
	return repository.GetUserByID(s.DB, id)
}

// UpdateUser updates an existing user
func (s *UserService) UpdateUser(user models.UserAccount) error {
	return repository.UpdateUser(s.DB, user)
}

// DeleteUser deletes a user
func (s *UserService) DeleteUser(user models.UserAccount) error {
	return repository.DeleteUser(s.DB, user)
}
