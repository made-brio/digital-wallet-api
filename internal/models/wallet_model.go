package models

type Wallet struct {
	ID      int     `json:"id"`      // Unique wallet ID
	UserID  int     `json:"user_id"` // Associated user ID
	Balance float64 `json:"balance"` // Current wallet balance
	Status  string  `json:"status"`  // Status: active, frozen
}
