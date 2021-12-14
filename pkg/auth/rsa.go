package auth

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/anish-sinha1/oidc-go/pkg/repo/models"
	"github.com/dgrijalva/jwt-go"
)

const (
	accessTokenPublicKeyPath  = "../../config/access_token/public_key.pem"
	accessTokenPrivateKeyPath = "../../config/access_token/private_key.pem"
)

type RSAKeyPair struct {
	PublicKey  []byte
	PrivateKey []byte
}

func NewRSAKeyPair() (Signer, error) {
	privateKey, err := ioutil.ReadFile(accessTokenPrivateKeyPath)
	if err != nil {
		return nil, err
	}
	publicKey, err := ioutil.ReadFile(accessTokenPublicKeyPath)
	if err != nil {
		return nil, err
	}
	return &RSAKeyPair{
		PublicKey:  publicKey,
		PrivateKey: privateKey,
	}, nil
}

func (k *RSAKeyPair) Issue(ttl time.Duration, sub string) (string, error) {
	privateKey, err := ioutil.ReadFile(accessTokenPrivateKeyPath)
	if err != nil {
		return "", err
	}
	jti := CreateJti()
	claims := models.AccessTokenClaims{}
	claims.Sub = sub
	claims.Iat = time.Now().Unix()
	claims.Exp = time.Now().Add(ttl).Unix()
	claims.Iss = "oidc-go"   // TODO: get from config
	claims.Aud = "react-app" // TODO: get from config
	claims.Jti = jti
	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, &claims).SignedString(privateKey) // Sign with private key
	if err != nil {
		return "", err
	}
	return token, nil
}

func (k *RSAKeyPair) Validate(token string) (models.Claims, error) {
	publicKey, err := ioutil.ReadFile(accessTokenPublicKeyPath)
	if err != nil {
		return nil, err
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
