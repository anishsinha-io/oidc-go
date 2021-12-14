package models

// StandardClaimsAddress defines the address JSON object returned by the standard
// claims at the /userinfo endpoint. Defined here: https://openid.net/specs/openid-connect-core-1_0.html#AddressClaim
type StandardClaimsAddress struct {
	Formatted     string `json:"formatted"`
	StreetAddress string `json:"street"`
	City          string `json:"city"`
	Locality      string `json:"locality"`
	Region        string `json:"region"`
	PostalCode    string `json:"postalCode"`
	Country       string `json:"country"`
}

// StandardClaims represents the OP server's response to the /userinfo enpdoint.
// Defined here: https://openid.net/specs/openid-connect-core-1_0.html#Claims
type StandardClaims struct {
	Sub                 string                 `json:"sub"`
	Name                string                 `json:"name"`
	GivenName           string                 `json:"given_name"`
	FamilyName          string                 `json:"family_name"`
	MiddleName          string                 `json:"middle_name"`
	Nickname            string                 `json:"nickname"`
	PreferredUsername   string                 `json:"preferred_username"`
	Profile             string                 `json:"profile"`
	Picture             string                 `json:"picture"`
	Website             string                 `json:"website"`
	Email               string                 `json:"email"`
	EmailVerified       bool                   `json:"email_verified"`
	Gender              string                 `json:"gender"`
	BirthDate           string                 `json:"birthdate"`
	ZoneInfo            string                 `json:"zoneinfo"`
	Locale              string                 `json:"locale"`
	PhoneNumber         string                 `json:"phone_number"`
	PhoneNumberVerified bool                   `json:"phone_number_verified"`
	Address             StandardClaimsAddress  `json:"address"`
	UpdatedAt           int                    `json:"updated_at"`
	CustomClaims        map[string]interface{} `json:"custom_claims,omitempty"`
}
