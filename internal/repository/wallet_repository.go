package repository

import (
	"database/sql"
	"digital-wallet-api/internal/models"
)

// CreateWallet inserts a new wallet into the database and updates the user's wallet_id
func CreateWallet(db *sql.DB, userId int) (err error) {
	// Insert a new wallet
	sqlQuery := `
		INSERT INTO wallet (user_id, balance, status) 
		VALUES ($1, $2, $3)
		RETURNING id` // Capture the newly created wallet ID

	var walletId int
	err = db.QueryRow(sqlQuery, userId, 0.0, "active").Scan(&walletId)
	if err != nil {
		return err
	}

	// Update the user's wallet_id
	updateUserQuery := `
		UPDATE user_account 
		SET wallet_id = $1 
		WHERE id = $2`

	_, err = db.Exec(updateUserQuery, walletId, userId)
	if err != nil {
		return err
	}

	return nil
}

// CheckBalance retrieves the wallet balance from the database
func CheckBalance(db *sql.DB, walletID int) (result models.Wallet, err error) {
	sqlQuery := "SELECT id, user_id, balance FROM wallet WHERE id = $1"

	// Execute query and map the result to the `result` object
	err = db.QueryRow(sqlQuery, walletID).Scan(
		&result.ID,
		&result.UserID,
		&result.Balance,
	)

	if err == sql.ErrNoRows {
		return result, nil // No wallet found, returning empty result
	} else if err != nil {
		return result, err
	}

	return result, nil
}

func GetInfoByWalletId(db *sql.DB, walletID int) (result models.Wallet, err error) {
	sqlQuery := "SELECT id, user_id, balance,status FROM wallet WHERE id = $1"

	// Execute query and map the result to the `result` object
	err = db.QueryRow(sqlQuery, walletID).Scan(
		&result.ID,
		&result.UserID,
		&result.Balance,
		&result.Status,
	)

	if err == sql.ErrNoRows {
		return result, nil // No wallet found, returning empty result
	} else if err != nil {
		return result, err
	}

	return result, nil
}

// TopUp adds funds to a wallet
func TopUp(db *sql.DB, payload models.TopUpRequest) (result models.Wallet, err error) {
	// Get existing wallet balance
	var currentBalance float64
	getBalanceQuery := "SELECT balance FROM wallet WHERE id = $1"
	err = db.QueryRow(getBalanceQuery, payload.WalletID).Scan(&currentBalance)
	if err != nil {
		if err == sql.ErrNoRows {
			return result, sql.ErrNoRows // Wallet not found
		}
		return result, err
	}

	// Add payload amount to the current balance
	newBalance := currentBalance + payload.Amount

	// Update wallet balance
	updateBalanceQuery := "UPDATE wallet SET balance = $1 WHERE id = $2"
	_, err = db.Exec(updateBalanceQuery, newBalance, payload.WalletID)
	if err != nil {
		return result, err
	}

	// Return updated wallet details
	result.ID = payload.WalletID
	result.Balance = newBalance
	return result, nil
}

// FreezeWallet updates the wallet status to "Frozen"
func FreezeWallet(db *sql.DB, walletID int) (err error) {
	sqlQuery := `UPDATE wallet SET status = 'Frozen' WHERE id = $1`
	_, err = db.Exec(sqlQuery, walletID)
	return err
}

// UnfreezeWallet updates the wallet status to "Active"
func UnfreezeWallet(db *sql.DB, walletID int) (err error) {
	sqlQuery := `UPDATE wallet SET status = 'Active' WHERE id = $1`
	_, err = db.Exec(sqlQuery, walletID)
	return err
}
