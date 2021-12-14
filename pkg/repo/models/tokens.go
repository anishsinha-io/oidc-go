package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Signer interface {
	Issue(time.Duration, string) (string, error)
	Validate(string) (Claims, error)
}

type Claims interface {
	Valid() error
}

type AccessTokenClaims struct {
	Aud      string    `json:"aud"`
	Exp      int64     `json:"exp"`
	Iat      int64     `json:"iat"`
	Iss      string    `json:"iss"`
	Sub      string    `json:"sub"`
	ClientId string    `json:"client_id"`
	Jti      uuid.UUID `json:"jti"`
}

type IdTokenClaims struct {
	Aud      string    `json:"aud"`
	Exp      int64     `json:"exp"`
	Iat      int64     `json:"iat"`
	Iss      string    `json:"iss"`
	Sub      string    `json:"sub"`
	ClientId string    `json:"client_id"`
	Jti      uuid.UUID `json:"jti"`
	Azp      string    `json:"azp"`
	Acr      string    `json:"acr"`
	Nonce    uuid.UUID `json:"nonce"`
	Amr      string    `json:"amr"`
}

type RefreshTokenClaims struct {
	Aud            string    `json:"aud"`
	Exp            int64     `json:"exp"`
	Iat            int64     `json:"iat"`
	Iss            string    `json:"iss"`
	Sub            string    `json:"sub"`
	ClientId       string    `json:"client_id"`
	Jti            uuid.UUID `json:"jti"`
	AccessTokenJti string    `json:"access_token_jti"`
}

func (t *AccessTokenClaims) Valid() error {
	if t.Exp < time.Now().Unix() {
		return errors.New("token expired")
	}
	return nil
}

func (t *IdTokenClaims) Valid() error {
	if t.Exp < time.Now().Unix() {
		return errors.New("token expired")
	}
	return nil
}

func (t *RefreshTokenClaims) Valid() error {
	if t.Exp < time.Now().Unix() {
		return errors.New("token expired")
	}
	return nil
}
