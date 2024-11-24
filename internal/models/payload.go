package models

import "github.com/golang-jwt/jwt"

type TransferRequest struct {
	FromWalletID int     `json:"from_wallet_id" binding:"required"`
	ToWalletID   int     `json:"to_wallet_id" binding:"required"`
	Amount       float64 `json:"amount" binding:"required"`
}

type TopUpRequest struct {
	WalletID int     `json:"id"`
	Amount   float64 `json:"amount" binding:"required"`
}

type UserLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserClaims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}
