package oauth

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/anish-sinha1/oidc-go/pkg/auth"
	"github.com/anish-sinha1/oidc-go/pkg/repo/models"
)

func ValidateToken(token string) (bool, error) {
	kp, err := auth.NewRSAKeyPair()
	if err != nil {
		return false, err
	}
	// validate signature
	claims, err := kp.Validate(token)
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

func IssueToken(ttl time.Duration, sub string) (string, error) {
	kp, err := auth.NewRSAKeyPair()
	if err != nil {
		return "", err
	}
	token, err := kp.Issue(ttl, sub)
	if err != nil {
		return "", err
	}
	return token, nil
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
