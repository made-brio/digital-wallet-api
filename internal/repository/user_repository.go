package repository

import (
	"database/sql"
	"digital-wallet-api/internal/models"
)

// GetAllUsers retrieves all users from the database
func GetAllUsers(db *sql.DB) (result []models.UserAccount, err error) {
	sqlQuery := "SELECT id, username, wallet_id FROM user_account"

	rows, err := db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Iterate through the rows and populate the result
	for rows.Next() {
		var user models.UserAccount
		err = rows.Scan(
			&user.ID,
			&user.Username,
			&user.WalletID,
		)
		if err != nil {
			return nil, err
		}

		result = append(result, user)
	}

	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

// GetUserByID retrieves a user by their ID from the database
func GetUserByID(db *sql.DB, id int) (result models.UserAccount, err error) {
	sqlQuery := "SELECT id, username, wallet_id FROM user_account WHERE id = $1"

	// Execute query and map the result to the `result` object
	err = db.QueryRow(sqlQuery, id).Scan(
		&result.ID,
		&result.Username,
		&result.WalletID,
	)

	// Check for no rows found
	if err == sql.ErrNoRows {
		return result, nil // No user found, returning empty result
	} else if err != nil {
		return result, err
	}

	return result, nil
}

// UpdateUser updates an existing user’s details in the database
func UpdateUser(db *sql.DB, user models.UserAccount) (err error) {
	sqlQuery := `UPDATE user_account SET username = $1 WHERE id = $2`
	_, err = db.Exec(sqlQuery, user.Username, user.ID)
	return err
}

// DeleteUser deletes a user from the database based on their ID
func DeleteUser(db *sql.DB, user models.UserAccount) (err error) {
	sqlQuery := "DELETE FROM user_account WHERE id = $1"
	_, err = db.Exec(sqlQuery, user.ID)
	return err
}
