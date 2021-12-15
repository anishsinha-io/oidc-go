package auth

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/anish-sinha1/oidc-go/pkg/repo/models"
	"github.com/golang-jwt/jwt/v4"
)

const (
	accessTokenPublicPathRSA  = "../../config/rsa/access_token/public_key.pem"
	accessTokenPrivatePathRSA = "../../config/rsa/access_token/private_key.pem"

	idTokenPublicPathRSA  = "../../config/rsa/id_token/public_key.pem"
	idTokenPrivatePathRSA = "../../config/rsa/id_token/private_key.pem"

	refreshTokenPublicPathRSA  = "../../config/rsa/refresh_token/public_key.pem"
	refreshTokenPrivatePathRSA = "../../config/rsa/refresh_token/private_key.pem"
)

// RSAKeyPair is a struct that holds the public and private keys for signing and verifying
// tokens signed with RS256
type RSAKeyPair struct {
	PublicKey  []byte
	PrivateKey []byte
}

type SigningKeychain struct {
	AccessTokenKeys  RSAKeyPair
	IdTokenKeys      RSAKeyPair
	RefreshTokenKeys RSAKeyPair
}

type Signer interface {
	Issue(time.Duration, string) (map[string]string, error)
	Validate(string, string) (models.Claims, error)
}

func CreateRSAKeyPair(publicKeyPath string, privateKeyPath string) (*RSAKeyPair, error) {
	publicKey, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		return nil, err
	}
	privateKey, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return nil, err
	}
	return &RSAKeyPair{
		PublicKey:  publicKey,
		PrivateKey: privateKey,
	}, nil
}

func CreateKeychain() (Signer, error) {
	accessTokenKeys, err := CreateRSAKeyPair(accessTokenPublicPathRSA, accessTokenPrivatePathRSA)
	if err != nil {
		return nil, err
	}
	idTokenKeys, err := CreateRSAKeyPair(idTokenPublicPathRSA, idTokenPrivatePathRSA)
	if err != nil {
		return nil, err
	}
	refreshTokenKeys, err := CreateRSAKeyPair(refreshTokenPublicPathRSA, refreshTokenPrivatePathRSA)
	if err != nil {
		return nil, err
	}
	return &SigningKeychain{
		AccessTokenKeys:  *accessTokenKeys,
		IdTokenKeys:      *idTokenKeys,
		RefreshTokenKeys: *refreshTokenKeys,
	}, nil
}

func CreateAccessToken(ttl time.Duration, sub string, signingKey []byte) *models.AccessTokenClaims {
	jti := CreateJti()
	return &models.AccessTokenClaims{
		Iss:   "https://localhost:8080",
		Sub:   sub,
		Aud:   "https://localhost:3000",
		Exp:   time.Now().Add(ttl).Unix(),
		Iat:   time.Now().Unix(),
		Scope: "openid",
		Jti:   jti,
	}
}

func CreateIdToken(ttl time.Duration, sub string, signingKey []byte) *models.IdTokenClaims {
	jti := CreateJti()
	nonce := CreateNonce()
	return &models.IdTokenClaims{
		Iss:   "https://localhost:8080",
		Sub:   sub,
		Aud:   "https://localhost:3000",
		Exp:   time.Now().Add(ttl).Unix(),
		Iat:   time.Now().Unix(),
		Jti:   jti,
		Nonce: nonce,
	}
}

func CreateRefreshToken(ttl time.Duration, atc *models.AccessTokenClaims, signingKey []byte) *models.RefreshTokenClaims {
	jti := CreateJti()
	return &models.RefreshTokenClaims{
		Iss:            "https://localhost:8080",
		Sub:            atc.Sub,
		Aud:            "https://localhost:3000",
		Exp:            time.Now().Add(ttl).Unix(),
		Iat:            time.Now().Unix(),
		Scope:          atc.Scope,
		AccessTokenJti: atc.Jti,
		Jti:            jti,
	}
}

func (k *RSAKeyPair) CheckSignature(token string, tokenType string) (models.Claims, error) {
	var publicKey []byte
	switch tokenType {
	case "access_token":
		pk, err := ioutil.ReadFile(accessTokenPublicPathRSA)
		if err != nil {
			return nil, err
		}
		publicKey = pk
	case "id_token":
		pk, err := ioutil.ReadFile(idTokenPublicPathRSA)
		if err != nil {
			return nil, err
		}
		publicKey = pk
	case "refresh_token":
		pk, err := ioutil.ReadFile(refreshTokenPublicPathRSA)
		if err != nil {
			return nil, err
		}
		publicKey = pk

	}

	res, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", jwtToken.Header["alg"])
		}
		return publicKey, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := res.Claims.(*models.AccessTokenClaims)
	if !ok || !res.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}

func (k *SigningKeychain) Issue(ttl time.Duration, sub string) (map[string]string, error) {
	ask, err := ioutil.ReadFile(accessTokenPrivatePathRSA)
	if err != nil {
		return nil, err
	}
	isk, err := ioutil.ReadFile(idTokenPrivatePathRSA)
	if err != nil {
		return nil, err
	}
	rsk, err := ioutil.ReadFile(refreshTokenPrivatePathRSA)
	if err != nil {
		return nil, err
	}
	atc := CreateAccessToken(ttl, sub, ask)
	if err != nil {
		return nil, err
	}
	itc := CreateIdToken(ttl, sub, isk)
	if err != nil {
		return nil, err
	}
	rtc := CreateRefreshToken(ttl, atc, rsk)
	if err != nil {
		return nil, err
	}
	signedAccessToken, err := jwt.NewWithClaims(jwt.SigningMethodRS256, atc).SignedString(ask)
	if err != nil {
		return nil, err
	}
	signedIdToken, err := jwt.NewWithClaims(jwt.SigningMethodRS256, itc).SignedString(isk)
	if err != nil {
		return nil, err
	}
	signedRefreshToken, err := jwt.NewWithClaims(jwt.SigningMethodRS256, rtc).SignedString(rsk)
	if err != nil {
		return nil, err
	}
	return map[string]string{
		"access_token":  signedAccessToken,
		"id_token":      signedIdToken,
		"refresh_token": signedRefreshToken,
	}, nil
}

func (k *SigningKeychain) Validate(token string, tokenType string) (models.Claims, error) {
	switch tokenType {
	case "access_token":
		claims, err := k.AccessTokenKeys.CheckSignature(token, tokenType)
		if err != nil {
			return nil, err
		}
		return claims, nil
	case "id_token":
		claims, err := k.IdTokenKeys.CheckSignature(token, tokenType)
		if err != nil {
			return nil, err
		}
		return claims, nil
	case "refresh_token":
		claims, err := k.RefreshTokenKeys.CheckSignature(token, tokenType)
		if err != nil {
			return nil, err
		}
		return claims, nil
	default:
		return nil, fmt.Errorf("invalid token type")
	}
}
