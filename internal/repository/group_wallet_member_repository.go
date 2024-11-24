package repository

import (
	"database/sql"
	"digital-wallet-api/internal/models"
)

// GetAllMembers retrieves all members of a specific group wallet
func GetAllMembers(db *sql.DB, groupWalletID int) (result []models.GroupWalletMember, err error) {
	sqlQuery := `
		SELECT id, group_wallet_id, user_id, name, contribution 
		FROM group_wallet_member 
		WHERE group_wallet_id = $1`

	rows, err := db.Query(sqlQuery, groupWalletID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var member models.GroupWalletMember
		err = rows.Scan(
			&member.ID,
			&member.GroupWalletID,
			&member.UserID,
			&member.Name,
			&member.Contribution,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, member)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

// AddMember adds a new member to a group wallet
func AddMember(db *sql.DB, member models.GroupWalletMember) (err error) {

	sqlQuery := `INSERT INTO group_wallet_member (group_wallet_id, user_id, name, contribution) VALUES ($1, $2, $3, $4)`
	_, err = db.Exec(sqlQuery, member.GroupWalletID, member.UserID, member.Name, member.Contribution)
	return err
}

// RemoveMember removes a member from a group wallet
func RemoveMember(db *sql.DB, groupWalletID, userID int) (err error) {
	sqlQuery := `DELETE FROM group_wallet_member WHERE group_wallet_id = $1 AND user_id = $2`
	_, err = db.Exec(sqlQuery, groupWalletID, userID)
	return err
}

// UpdateMemberContribution updates the contribution of a member in a group wallet
func UpdateMemberContribution(db *sql.DB, member models.GroupWalletMember) (err error) {

	sqlQuery := `UPDATE group_wallet_member SET contribution = $1 WHERE group_wallet_id = $2 AND user_id = $3`
	_, err = db.Exec(sqlQuery, member.Contribution, member.GroupWalletID, member.UserID)
	return err
}
