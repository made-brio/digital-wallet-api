package models

type UserAccount struct {
	ID       int    `json:"id"`        // Unique identifier for the user
	Username string `json:"username"`  // User's name
	WalletID int    `json:"wallet_id"` // Linked wallet ID
	Password string `json:"password"`  // User's password (stored securely)
}
