package domain

import (
	"time"
)

type TokenRequest struct {
	GrantType    string `json:"grant_type" binding:"required,oneof=authorization_code"`
	Code         string `json:"code" binding:"required"`
	RedirectUri  string `json:"redirect_uri" binding:"required"`
	ClientId     string `json:"client_id" binding:"required"`
	CodeVerifier string `json:"code_verifier" binding:"required"`
}

type AccessToken struct {
	Token     string    `gorm:"primaryKey"`
	UserId    string    `gorm:"not null"`
	ClientId  string    `gorm:"not null"`
	Scope     string    `gorm:"not null"`
	ExpiredAt time.Time `gorm:"not null"`
}

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiredIn   int    `json:"expired_in"`
	Scope       string `json:"scope"`
}
