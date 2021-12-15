package oauth

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/anish-sinha1/oidc-go/pkg/auth"
	"github.com/anish-sinha1/oidc-go/pkg/repo/models"
)

// ValidateToken validates a token and returns a boolean. It can be passed any string and
//  any signing method, allowing it to be used with any signing algorithm
func ValidateToken(token string, signingMethod string, tokenType string) (bool, error) {
	var signer auth.Signer
	switch signingMethod {
	case "RS256":
		kp, err := auth.CreateKeychain()
		if err != nil {
			return false, err
		}
		signer = kp
	case "HS256":
		// implement HS256
	case "ES256":
		// implement ES256
	}
	// validate signature
	claims, err := signer.Validate(token, tokenType)
	if err != nil {
		return false, err
	}

	// validate expiry
	err = claims.Valid()
	if err != nil {
		return false, err
	}
	return true, nil
}

func IssueTokens(ttl time.Duration, sub string, token string) (map[string]string, error) {
	keychain, err := auth.CreateKeychain()
	if err != nil {
		return nil, err
	}
	tokens, err := keychain.Issue(ttl, sub)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}

func ValidateUriParams(qs models.AuthorizeQs) (bool, error) {
	if openid := strings.Contains(qs.Scope, "openid"); !openid {
		return false, errors.New("invalid_request")
	}
	if qs.ResponseType != "code" {
		return false, errors.New("invalid_grant")
	}
	if qs.RedirectUri != os.Getenv("REDIRECT_URI") {
		return false, errors.New("invalid redirect_uri")
	}
	if qs.ClientId != "test_client_id" {
		return false, errors.New("invalid_client")
	}
	if qs.State == "" {
		return false, errors.New("invalid_request")
	}
	return true, nil
}
