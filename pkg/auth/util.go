package auth

import (
	"github.com/google/uuid"
)

// CreateNonce is for generating nonces for Id Tokens
func CreateNonce() uuid.UUID {
	return uuid.New()
}

// CreateNewAuthCode is for generating authorization codes in the authorization code flow
func CreateNewAuthCode() uuid.UUID {
	return uuid.New()
}

func CreateJti() uuid.UUID {
	return uuid.New()
}
