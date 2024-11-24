package models

import "time"

type GroupWallet struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Goal      float64   `json:"goal"`
	CreatedBy int       `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}
