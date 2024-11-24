package repository

import (
	"database/sql"
	"digital-wallet-api/internal/models"
	"time"
)

// GetGroupWalletByID retrieves a specific group wallet by its ID
func GetGroupWalletByID(db *sql.DB, id int) (result models.GroupWallet, err error) {
	sqlQuery := "SELECT id, name, goal, created_by, created_at FROM group_wallet WHERE id = $1"

	err = db.QueryRow(sqlQuery, id).Scan(
		&result.ID,
		&result.Name,
		&result.Goal,
		&result.CreatedBy,
		&result.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return result, nil // No group wallet found
	} else if err != nil {
		return result, err
	}

	return result, nil
}

// CreateGroupWallet creates a new group wallet
func CreateGroupWallet(db *sql.DB, groupWallet models.GroupWallet) (err error) {
	sqlQuery := `INSERT INTO group_wallet (name, goal, created_by, created_at) VALUES ($1, $2, $3, $4)`
	_, err = db.Exec(sqlQuery, groupWallet.Name, groupWallet.Goal, groupWallet.CreatedBy, time.Now())
	return err
}

// UpdateGroupWallet updates an existing group wallet
func UpdateGroupWalletGoal(db *sql.DB, groupWallet models.GroupWallet) (err error) {
	sqlQuery := `UPDATE group_wallet SET goal = $1 WHERE id = $2`
	_, err = db.Exec(sqlQuery, groupWallet.Goal, groupWallet.ID)
	return err
}

// DeleteGroupWallet deletes a group wallet by its ID
func DeleteGroupWallet(db *sql.DB, id int) (err error) {
	sqlQuery := "DELETE FROM group_wallet WHERE id = $1"
	_, err = db.Exec(sqlQuery, id)
	return err
}
