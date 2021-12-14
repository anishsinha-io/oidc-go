package auth

import (
	"time"

	"github.com/anish-sinha1/oidc-go/pkg/repo/models"
)

type Signer interface {
	Issue(time.Duration, string) (string, error)
	Validate(string) (models.Claims, error)
}
