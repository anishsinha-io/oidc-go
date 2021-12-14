package app

type AppConfig struct {
}

type OpenIdConfig struct {
	Issuer               string
	ClientId             string
	AllowedRedirectUri   string
	AllowedScopes        string
	AllowedResponseTypes string
}
