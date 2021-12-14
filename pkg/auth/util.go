package auth

import (
	"encoding/base64"
	"math/rand"

	"github.com/google/uuid"
)

// CreateNonce is for generating nonces for Id Tokens
func CreateNonce() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return base64.StdEncoding.EncodeToString(bytes)
}

// CreateNewAuthCode is for generating authorization codes in the authorization code flow
func CreateNewAuthCode() uuid.UUID {
	return uuid.New()
}

func CreateJti() uuid.UUID {
	return uuid.New()
}
