package models

type GroupWalletMember struct {
	ID            int     `json:"id"`
	GroupWalletID int     `json:"group_wallet_id"`
	UserID        int     `json:"user_id"`
	Name          string  `json:"name"`
	Contribution  float64 `json:"contribution"`
}
