package main

import (
	"os"

	"github.com/anish-sinha1/oidc-go/pkg/app"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var OpenIdConfig app.OpenIdConfig

func main() {
	OpenIdConfig = app.OpenIdConfig{
		Issuer:               os.Getenv("ISSUER"),
		ClientId:             os.Getenv("CLIENT_ID"),
		AllowedRedirectUri:   os.Getenv("ALLOWED_REDIRECT_URI"),
		AllowedScopes:        os.Getenv("ALLOWED_SCOPES"),
		AllowedResponseTypes: os.Getenv("ALLOWED_RESPONSE_TYPES"),
	}
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	app.Router = gin.Default()
	app.InitializeRoutes()
	app.Router.Run(":8080")
}
