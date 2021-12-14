package models

import (
	"github.com/google/uuid"
)

type Session struct {
	Id                int64     `json:"id"`
	AccessToken       string    `json:"access_token"`
	RefreshToken      string    `json:"refresh_token"`
	AuthorizationCode uuid.UUID `json:"authorization_code"`
	State             string    `json:"state"`
	Nonce             string    `json:"nonce"`
}
