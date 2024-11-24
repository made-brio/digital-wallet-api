package service

import (
	"database/sql"
	"digital-wallet-api/internal/models"
	"digital-wallet-api/internal/repository"
)

// GroupWalletMemberService defines the service layer for group wallet member operations
type GroupWalletMemberService struct {
	DB *sql.DB
}

// NewGroupWalletMemberService creates a new instance of GroupWalletMemberService
func NewGroupWalletMemberService(db *sql.DB) *GroupWalletMemberService {
	return &GroupWalletMemberService{DB: db}
}

// AddMember adds a new member to a group wallet
func (gwms *GroupWalletMemberService) AddMember(member models.GroupWalletMember) error {
	return repository.AddMember(gwms.DB, member)
}

// GetMembers retrieves all members of a group wallet
func (gwms *GroupWalletMemberService) GetAllMembers(walletID int) ([]models.GroupWalletMember, error) {
	return repository.GetAllMembers(gwms.DB, walletID)
}

// UpdateMemberContribution updates the contribution of a specific member
func (gwms *GroupWalletMemberService) UpdateMemberContribution(member models.GroupWalletMember) error {
	return repository.UpdateMemberContribution(gwms.DB, member)
}

// DeleteMember removes a member from the group wallet
func (gwms *GroupWalletMemberService) RemoveMember(groupWalletID, memberID int) error {
	return repository.RemoveMember(gwms.DB, groupWalletID, memberID)
}
