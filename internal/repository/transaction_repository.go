package repository

import (
	"database/sql"
	"digital-wallet-api/internal/models"
	"strconv"
)

// Transfer handles fund transfers between wallets
func Transfer(db *sql.DB, transfer models.TransferRequest) (err error) {
	// Begin transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// Deduct from sender's wallet
	deductQuery := "UPDATE wallet SET balance = balance - $1 WHERE id = $2 AND balance >= $1"
	res, err := tx.Exec(deductQuery, transfer.Amount, transfer.FromWalletID)
	if err != nil {
		tx.Rollback()
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		tx.Rollback()
		return sql.ErrNoRows // Insufficient funds or wallet not found
	}

	// Add a debit transaction for the sender
	debitTransactionQuery := `
		INSERT INTO transaction (wallet_id, type, amount, other_party, description, timestamp) 
		VALUES ($1, 'expense', $2, $3, $4, NOW())`
	_, err = tx.Exec(debitTransactionQuery, transfer.FromWalletID, transfer.Amount, transfer.ToWalletID, "Fund transfer to wallet "+strconv.Itoa(transfer.ToWalletID))
	if err != nil {
		tx.Rollback()
		return err
	}

	// Add to receiver's wallet
	addQuery := "UPDATE wallet SET balance = balance + $1 WHERE id = $2"
	_, err = tx.Exec(addQuery, transfer.Amount, transfer.ToWalletID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Add a credit transaction for the receiver
	creditTransactionQuery := `
		INSERT INTO transaction (wallet_id, type, amount, other_party, description, timestamp) 
		VALUES ($1, 'income', $2, $3, $4, NOW())`
	_, err = tx.Exec(creditTransactionQuery, transfer.ToWalletID, transfer.Amount, transfer.FromWalletID, "Fund transfer from wallet "+strconv.Itoa(transfer.FromWalletID))
	if err != nil {
		tx.Rollback()
		return err
	}

	// Commit transaction
	err = tx.Commit()
	return err
}

// IncomeHistory retrieves all income (credit) transactions for a wallet
func IncomeHistory(db *sql.DB, walletID int) (result []models.Transaction, err error) {
	sqlQuery := `SELECT id, wallet_id, type, amount, other_party, description, timestamp 
                 FROM transaction WHERE wallet_id = $1 AND type = 'income'`

	rows, err := db.Query(sqlQuery, walletID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction models.Transaction
		err = rows.Scan(&transaction.ID, &transaction.WalletID, &transaction.Type, &transaction.Amount,
			&transaction.OtherParty, &transaction.Description, &transaction.Timestamp)
		if err != nil {
			return nil, err
		}
		result = append(result, transaction)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

// ExpenseHistory retrieves all expense (debit) transactions for a wallet
func ExpenseHistory(db *sql.DB, walletID int) (result []models.Transaction, err error) {
	sqlQuery := `SELECT id, wallet_id, type, amount, other_party, description, timestamp 
                 FROM transaction WHERE wallet_id = $1 AND type = 'expense'`

	rows, err := db.Query(sqlQuery, walletID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction models.Transaction
		err = rows.Scan(&transaction.ID, &transaction.WalletID, &transaction.Type, &transaction.Amount,
			&transaction.OtherParty, &transaction.Description, &transaction.Timestamp)
		if err != nil {
			return nil, err
		}
		result = append(result, transaction)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

// TransactionsHistory retrieves all transactions for a wallet
func TransactionsHistory(db *sql.DB, walletID int) (result []models.Transaction, err error) {
	sqlQuery := `SELECT id, wallet_id, type, amount, other_party, description, timestamp 
                 FROM transaction WHERE wallet_id = $1`

	rows, err := db.Query(sqlQuery, walletID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction models.Transaction
		err = rows.Scan(&transaction.ID, &transaction.WalletID, &transaction.Type, &transaction.Amount,
			&transaction.OtherParty, &transaction.Description, &transaction.Timestamp)
		if err != nil {
			return nil, err
		}
		result = append(result, transaction)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
