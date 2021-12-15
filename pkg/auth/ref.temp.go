package auth

// // NewRSAKeyPair generates a new RSA key pair given a token type, which can be access_token,
// // id_token, or refresh_token
// func NewRSAKeyPair(tokenType string) (Signer, error) {
// 	var privateKey, publicKey []byte
// 	switch tokenType {
// 	case "access_token":
// 		pvk, err := ioutil.ReadFile(accessTokenPrivatePathRSA)
// 		if err != nil {
// 			return nil, err
// 		}
// 		privateKey = pvk
// 		pbk, err := ioutil.ReadFile(accessTokenPublicPathRSA)
// 		if err != nil {
// 			return nil, err
// 		}
// 		publicKey = pbk
// 	case "id_token":
// 		pvk, err := ioutil.ReadFile(idTokenPrivatePathRSA)
// 		if err != nil {
// 			return nil, err
// 		}
// 		privateKey = pvk
// 		pbk, err := ioutil.ReadFile(idTokenPublicPathRSA)
// 		if err != nil {
// 			return nil, err
// 		}
// 		publicKey = pbk
// 	case "refresh_token":
// 		pvk, err := ioutil.ReadFile(refreshTokenPrivatePathRSA)
// 		if err != nil {
// 			return nil, err
// 		}
// 		privateKey = pvk
// 		pbk, err := ioutil.ReadFile(refreshTokenPublicPathRSA)
// 		if err != nil {
// 			return nil, err
// 		}
// 		publicKey = pbk
// 	default:
// 		return nil, fmt.Errorf("invalid token type")
// 	}

// 	return &RSAKeyPair{
// 		PublicKey:  publicKey,
// 		PrivateKey: privateKey,
// 	}, nil
// }

// func NewRSAKeychain() (SigningKeychain, error) {
// 	accessTokenPrivateKey, err := ioutil.ReadFile(accessTokenPrivatePathRSA)
// 	if err != nil {
// 		return SigningKeychain{}, err
// 	}
// 	accessTokenPublicKey, err := ioutil.ReadFile(accessTokenPublicPathRSA)
// 	if err != nil {
// 		return SigningKeychain{}, err
// 	}
// 	idTokenPrivateKey, err := ioutil.ReadFile(idTokenPrivatePathRSA)
// 	if err != nil {
// 		return SigningKeychain{}, err
// 	}
// 	idTokenPublicKey, err := ioutil.ReadFile(idTokenPublicPathRSA)
// 	if err != nil {
// 		return SigningKeychain{}, err
// 	}
// 	refreshTokenPrivateKey, err := ioutil.ReadFile(refreshTokenPrivatePathRSA)
// 	if err != nil {
// 		return SigningKeychain{}, err
// 	}

// 	refreshTokenPublicKey, err := ioutil.ReadFile(refreshTokenPublicPathRSA)
// 	if err != nil {
// 		return SigningKeychain{}, err
// 	}
// 	return SigningKeychain{
// 		AccessTokenKeys: RSAKeyPair{
// 			PublicKey:  accessTokenPublicKey,
// 			PrivateKey: accessTokenPrivateKey,
// 		},
// 		IdTokenKeys: RSAKeyPair{
// 			PublicKey:  idTokenPublicKey,
// 			PrivateKey: idTokenPrivateKey,
// 		},
// 		RefreshTokenKeys: RSAKeyPair{
// 			PublicKey:  refreshTokenPublicKey,
// 			PrivateKey: refreshTokenPrivateKey,
// 		},
// }

// // Issue creates and signs a new JWT. If the tokenType is access_token it will also generate a corresponding
// // refresh token with the right AccessTokenJti. If the token_type is id_token it will only generate an id token.
// func (k *RSAKeyPair) Issue(ttl time.Duration, sub string) (map[string]string, error) {
// 	privateKey, err := ioutil.ReadFile(accessTokenPrivateKeyPath)
// 	if err != nil {
// 		return "", err
// 	}
// 	jti := CreateJti()
// 	claims := models.AccessTokenClaims{}
// 	claims.Sub = sub
// 	claims.Iat = time.Now().Unix()
// 	claims.Exp = time.Now().Add(ttl).Unix()
// 	claims.Iss = "oidc-go"   // TODO: get from config
// 	claims.Aud = "react-app" // TODO: get from config
// 	claims.Jti = jti
// 	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, &claims).SignedString(privateKey) // Sign with private key
// 	if err != nil {
// 		return "", err
// 	}
// 	return token, nil
// }

// func (k *RSAKeyPair) Validate(token string) (models.Claims, error) {
// 	publicKey, err := ioutil.ReadFile(accessTokenPublicPathRSA)
// 	if err != nil {
// 		return nil, err
// 	}
// 	res, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
// 		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", jwtToken.Header["alg"])
// 		}
// 		return publicKey, nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	claims, ok := res.Claims.(*models.AccessTokenClaims)
// 	if !ok || !res.Valid {
// 		return nil, fmt.Errorf("invalid token")
// 	}
// 	return claims, nil
// }
