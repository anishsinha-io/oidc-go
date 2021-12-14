package models

type IntrospectionRequest struct {
	Token string `json:"token"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignupRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type AuthorizeQs struct {
	ClientId     string `form:"client_id"`
	RedirectUri  string `form:"redirect_uri"`
	ResponseType string `form:"response_type"`
	Scope        string `form:"scope"`
	State        string `form:"state"`
	Nonce        string `form:"nonce,omitempty"`
	Display      string `form:"display,omitempty"`
	Prompt       string `form:"prompt,omitempty"`
	MaxAge       string `form:"max_age,omitempty"`
	UiLocales    string `form:"ui_locales,omitempty"`
	IdTokenHint  string `form:"id_token_hint,omitempty"`
	LoginHint    string `form:"login_hint,omitempty"`
	AcrValues    string `form:"acr_values,omitempty"`
}
